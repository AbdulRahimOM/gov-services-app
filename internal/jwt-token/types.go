package jwttoken

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ExtraPurposeInfo struct {
	Purpose    string
	ExpiryTime time.Time
}

func (e *ExtraPurposeInfo) isCustomClaimType() {}

func (e *ExtraPurposeInfo) SetContext(c *gin.Context) {
	c.Set("purpose", e.Purpose)
	c.Set("purpose-expiry-time", e.ExpiryTime)
	fmt.Println("e.Purpose:", e.Purpose)
}

func (e *ExtraPurposeInfo) getWithLabel() additionalInfoWithLabel {
	fmt.Println("==--=--=-=-=-=-=-=getting with label")
	return additionalInfoWithLabel{
		Label: "ExtraPurposeInfo",
		Info:  e,
	}
}
func (e *ExtraPurposeInfo) getFromInfoMap(infoMap map[string]interface{}) (*ExtraPurposeInfo, error) {
	const layout = "2006-01-02T15:04:05.999999-07:00"
	infoMap2 := infoMap["Info"].(map[string]interface{})
	parsedTime, err := time.Parse(layout, infoMap2["ExpiryTime"].(string))
	if err != nil {
		fmt.Println("error in parsing time:", err)
		return nil, err
	}
	extraPurposeInfo := ExtraPurposeInfo{
		Purpose:    infoMap2["Purpose"].(string),
		ExpiryTime: parsedTime,
	}
	return &extraPurposeInfo, nil
}
