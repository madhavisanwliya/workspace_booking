package controller

import (
	"strconv"
	"workspace_booking/config"
	"workspace_booking/mailer"
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// CreateBooking handler
func CreateBooking(c *fiber.Ctx) error {
	timingParams := new(model.BookingTiming)

	c.BodyParser(timingParams)

	fromDatetTime, toDateTime := model.BookingTimestamp(timingParams)

	workspaceParams := new(model.Booking)

	if err := c.BodyParser(workspaceParams); err != nil {
		return utility.ErrResponse(c, "Error in body parsing", 400, err)
	}

	workspaceParams.FromDateTime = fromDatetTime
	workspaceParams.ToDateTime = toDateTime

	err := workspaceParams.InsertBooking()

	if err != nil {
		return utility.ErrResponse(c, "Error in creation", 500, err)
	}

	err = model.BulkInsertBookingParticipant(workspaceParams)

	err = model.BulkInsertBookingWorkspace(workspaceParams, timingParams)

	if err != nil {
		return utility.ErrResponse(c, "Error in creating participants", 500, err)
	}

	if workspaceParams.Id != 0 {
		go mailer.BookingMailer(workspaceParams.Id, false)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Booking successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in response", 500, err)
	}
	return nil
}

func GetAvailableBookingSpace(c *fiber.Ctx) error {
	reqFloorId := c.Query("floor_id")
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	reqbuildingId := c.Query("building_id")
	floorId, err := strconv.Atoi(reqFloorId)
	buildingId, err := strconv.Atoi(reqbuildingId)
	userIds := c.Query("user_ids")
	purpose := c.Query("purpose")

	if err != nil {
		return utility.ErrResponse(c, "Error in string convertion", 500, err)
	}
	timingParams := new(model.BookingTiming)
	timingParams.FromDate = fromDate
	timingParams.ToDate = toDate
	timingParams.StartTime = startTime
	timingParams.EndTime = endTime

	fromDatetTime, toDateTime := model.BookingTimestamp(timingParams)
	city := model.GetCityByFloorId(buildingId)
	// getting booking worksapcesspace
	availableWorkSpace, err := model.GetAvailableBookingSpace(floorId, fromDatetTime, toDateTime)
	availableWorkSpace.FromDate = fromDate
	availableWorkSpace.ToDate = toDate
	availableWorkSpace.StartTime = startTime
	availableWorkSpace.EndTime = endTime
	availableWorkSpace.CityId = int(city.Id)
	availableWorkSpace.CityName = city.Name
	availableWorkSpace.Purpose = purpose

	if err := c.JSON(&fiber.Map{
		"success":  true,
		"data":     availableWorkSpace,
		"user_ids": userIds,
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting available booking", 500, err)
	}
	return nil
}

func WorkSpacesDetails(c *fiber.Ctx) error {

	workspaceDetails := model.GetAllDetails()
	if err := c.JSON(&fiber.Map{
		"success":           true,
		"workspace_details": workspaceDetails,
		"message":           "All workspace details returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting workspace details", 500, err)
	}
	return nil
}

func MyBookingDetails(c *fiber.Ctx) error {
	auth, err := config.GetAuthDetails(c)
	if err != nil {
		return utility.ErrResponse(c, "Error in getting buildings", 500, err)
	}

	var userId int

	userId, _ = strconv.Atoi(auth.UserID)

	workspaceDetails := model.GetMyBookingDetails(true, userId)
	pastBookingDetails := model.GetMyBookingDetails(false, userId)
	if err := c.JSON(&fiber.Map{
		"success":                  true,
		"upcoming_booking_details": workspaceDetails,
		"past_booking_details":     pastBookingDetails,
		"message":                  "All My bookings returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting My bookings", 500, err)
	}
	return nil
}

func CanceledBookings(c *fiber.Ctx) error {
	id := c.Params("id")
	i, e := strconv.ParseInt(id, 0, 16)
	if e != nil {
		return c.Status(400).SendString(e.Error())
	}
	v := model.Booking{Id: int16(i)}
	err := model.CancelBooking(v)
	if err == nil {
		go mailer.CancelMailer(v.Id)
	}
	if err != nil {
		return utility.ErrResponse(c, "Error in  canceling", 400, err)
	}
	return c.JSON(fiber.Map{
		"message": "Meeting has been canceled",
	})
}
