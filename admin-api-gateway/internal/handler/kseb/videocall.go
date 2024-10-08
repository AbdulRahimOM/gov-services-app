package ksebhanlder

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/pion/webrtc/v3"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	w "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/webrtc"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
)

// VideoCallRoom
func (kseb *KSEBHandler) VideoCallRoom(c *fiber.Ctx) error {
	complaintId, err := gateway.HandleGetUrlParamsInt32Fiber(c, "complaintId")
	if err != nil {
		return err
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}

	_ = createOrGetRoom(complaintId)
	return c.Status(http.StatusOK).Render("peer", fiber.Map{
		"Content":           "peer",
		"RoomWebsocketAddr": fmt.Sprintf("%s://%s/kseb/ouser/videocall/room/%d/websocket", ws, config.EnvValues.UserHost, complaintId),
		"RoomLink":          fmt.Sprintf("%s://%s/kseb/oadmin/videocall/room/%d", c.Protocol(), c.Hostname(), complaintId),
		"Type":              "room",
	}, "layouts/main")
}

func (kseb *KSEBHandler) RoomWebsocket(c *websocket.Conn) {
	complaintIdStr := c.Params("complaintId")
	complaintId, err := strconv.Atoi(complaintIdStr)
	if err != nil {
		fmt.Println("error while looking for complaintId url param")
		return
	}
	room := createOrGetRoom(int32(complaintId))
	w.RoomConn(c, room.Peers)
}

func createOrGetRoom(complaintId int32) *w.Room {
	w.RoomsLock.Lock()
	defer w.RoomsLock.Unlock()

	if room := w.Rooms[complaintId]; room != nil {
		return room
	}

	p := &w.Peers{}
	p.TrackLocals = make(map[string]*webrtc.TrackLocalStaticRTP)
	room := &w.Room{
		Peers: p,
	}
	w.Rooms[complaintId] = room
	return room
}
