package twilioOTP

import (
	"fmt"

	"github.com/twilio/twilio-go"

	// twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioClient struct {
	bypassMode bool
	client     *twilio.RestClient
	serviceSid string
}

func NewTwilioClient(accountSid, authToken, serviceSid string, byPassTwilio bool) *TwilioClient {
	return &TwilioClient{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		}),
		serviceSid: serviceSid,
		bypassMode: byPassTwilio,
	}
}

// func (tc *TwilioClient) SendOtp(phone string) error {
// 	fmt.Println("Sending OTP")
// 	fmt.Println("phone", phone)
// 	params := &twilioApi.CreateMessageParams{}
// 	params.SetTo(phone)

// 	params.SetBody("Hello from Go!")
// 	resp, err := tc.client.Api.CreateMessage(params)
// 	if err != nil {
// 		return fmt.Errorf("error sending otp message, error: %v", err)
// 	} else {
// 		response, _ := json.Marshal(*resp)
// 		fmt.Println("Response: " + string(response))
// 		if resp.Status != nil { //?
// 			switch *resp.Status {
// 			case "pending":
// 				fmt.Println("OTP had been sent already")
// 			case "max_attempts_reached":
// 				fmt.Println("Max attempts reached")
// 			case "failed":
// 				fmt.Println("Failed")
// 				//need proper documentation reading 	//to be updated #need update
// 			}

// 			if *resp.Status == "pending" {
// 				// fmt.Println("OTP had been sent already")
// 				fmt.Print("") //blank code //implement later if needed
// 			}
// 		}
// 		return nil
// 	}
// }

// func (tc *TwilioClient) VerifyOtp(phone string, otp string) (bool, error) {
// 	fmt.Println("Verifying OTP")
// 	return true, nil
// fmt.Println("tc.serviceSid", tc.serviceSid)
// params := &twilioApi.CreateMessageParams{}

// params.
// // params := &verify.CreateVerificationCheckParams{}
// // params.SetTo(phone)
// // params.SetCode(otp)

// resp, err := tc.client.VerifyV2.CreateVerificationCheck(tc.serviceSid, params)
// if err != nil {
// 	fmt.Println(err.Error())
// 	return false, err
// } else {
// 	if resp.Status != nil {
// 		switch *resp.Status {
// 		//pending, approved, canceled, max_attempts_reached, deleted, failed or expired.
// 		case "approved":
// 			fmt.Println("OTP verification approved")
// 			return true, nil
// 		case "pending":

// 			fmt.Println("OTP verification pending. OTP already sent")
// 		case "denied":
// 			fmt.Println("OTP verification denied")
// 		}
// 		fmt.Println(*resp.Status)
// 	} else {
// 		fmt.Println(resp.Status)
// 	}
// 	return true, nil
// }
// }

func (tc *TwilioClient) SendOtp(phone string) error {
	if tc.bypassMode{
		fmt.Println("ByPass mode is turned on in environment. Skipping sending of otp")
		return nil
	}
	fmt.Println("Sending OTP")
	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")
	resp, err := tc.client.VerifyV2.CreateVerification(tc.serviceSid, params)
	if err != nil {
		return err
	} else {
		if resp.Status != nil { //?
			switch *resp.Status {
			case "pending":
				fmt.Println("OTP had been sent already")
			case "max_attempts_reached":
				fmt.Println("Max attempts reached")
			case "failed":
				fmt.Println("Failed")
				//need proper documentation reading 	//to be updated #need update
			}

			if *resp.Status == "pending" {
				// fmt.Println("OTP had been sent already")
				fmt.Print("") //blank code //implement later if needed
			}
		}
		return nil
	}
}

func (tc *TwilioClient) VerifyOtp(phone string, otp string) (bool, error) {
	// fmt.Println("phone", phone)
	// fmt.Println("otp", otp)
	if tc.bypassMode {
		fmt.Println("ByPass mode is turned on in environment. Skipping sending of otp")
		return true,nil
	}
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	// params.SetVerificationSid(tc.serviceSid)
	params.SetCode(otp)

	resp, err := tc.client.VerifyV2.CreateVerificationCheck(tc.serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	} else {
		if resp.Status != nil {
			switch *resp.Status {
			//pending, approved, canceled, max_attempts_reached, deleted, failed or expired.
			case "approved":
				fmt.Println("OTP verification approved")
				return true, nil
			default:
				fmt.Println("OTP verification failed: ", *resp.Status)
				return false, fmt.Errorf("OTP verification failed")
			}
		} else {
			fmt.Println("resp status:", resp.Status)
		}
		return true, nil
	}
}
