package service

import (
	"context"
	"testing"
	"time"

	"backend-bootcamp-assignment-2024/internal/mapper"
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	mock_service "backend-bootcamp-assignment-2024/internal/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlatService_GetFlats(t *testing.T) {
	t.Parallel()
	var (
		ctx         = context.Background()
		houseId     = 1
		isModerator = true
		resp        = []entity.Flat{
			{
				Id:      1,
				HouseId: 1,
				Price:   1000,
				Rooms:   1,
				Status:  entity.FLATSTATUS_APPROVED,
			},
			{
				Id:      2,
				HouseId: 1,
				Price:   1000,
				Rooms:   1,
				Status:  entity.FLATSTATUS_APPROVED,
			},
			{
				Id:      3,
				HouseId: 1,
				Price:   1000,
				Rooms:   1,
				Status:  entity.FLATSTATUS_APPROVED,
			},
		}
		flats = []response.Flat{}
	)
	for _, f := range resp {
		flats = append(flats, *mapper.FlatEntityToFlatResponse(&f))
	}

	t.Run("success get flats", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		tm := mock_service.NewMockTransactionManager(ctrl)
		fr := mock_service.NewMockFlatRepository(ctrl)
		hr := mock_service.NewMockHouseRepository(ctrl)
		sr := mock_service.NewMockSubscriberRepository(ctrl)
		sender := mock_service.NewMockSender(ctrl)
		house := NewHouseService(hr, tm)
		subscriber := NewSubscriberService(sr, tm, sender)
		flat := NewFlatService(fr, house, subscriber, tm)

		fr.EXPECT().GetFlatsByHouseId(gomock.Any(), int32(houseId), isModerator).Return(resp, nil)
		res, err := flat.GetFlats(ctx, int32(houseId), isModerator)
		require.NoError(t, err)
		assert.Equal(t, flats, res)
	})
}

func TestFlatService_AddFlat(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		req = request.CreateFlat{
			HouseId: 1,
			Price:   1000,
			Rooms:   nil,
		}
		room = int32(1)
		_    = request.CreateFlat{
			HouseId: 1,
			Price:   1000,
			Rooms:   &room,
			Status:  entity.FLATSTATUS_CREATED,
		}
		ent = entity.Flat{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   1,
			Status:  entity.FLATSTATUS_CREATED,
		}
	)
	t.Run("success add flat", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		tm := mock_service.NewMockTransactionManager(ctrl)
		fr := mock_service.NewMockFlatRepository(ctrl)
		hr := mock_service.NewMockHouseRepository(ctrl)
		sr := mock_service.NewMockSubscriberRepository(ctrl)
		sender := mock_service.NewMockSender(ctrl)
		house := NewHouseService(hr, tm)
		subscriber := NewSubscriberService(sr, tm, sender)
		flat := NewFlatService(fr, house, subscriber, tm)

		tm.EXPECT().Tx(gomock.Any(), gomock.Any()).Return(nil)
		//fr.EXPECT().CreateFlat(gomock.Any(), req2).Return(&ent, nil)
		//hr.EXPECT().UpdateHouse(gomock.Any(), ent.HouseId, gomock.Any()).Return(nil)
		sr.EXPECT().GetSubsByHouseId(gomock.Any(), ent.HouseId)
		_, err := flat.AddFlat(ctx, req)
		time.Sleep(time.Duration(1) * time.Second)
		require.NoError(t, err)
	})
}

func TestFlatService_UpdateFlat(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		req = request.UpdateFlat{
			Id:     1,
			Status: entity.FLATSTATUS_ON_MODERATION,
		}
		ent2 = entity.Flat{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   1,
			Status:  entity.FLATSTATUS_ON_MODERATION,
		}
		ent1 = entity.Flat{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   1,
			Status:  entity.FLATSTATUS_CREATED,
		}
	)
	t.Run("success update flat", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		tm := mock_service.NewMockTransactionManager(ctrl)
		fr := mock_service.NewMockFlatRepository(ctrl)
		hr := mock_service.NewMockHouseRepository(ctrl)
		sr := mock_service.NewMockSubscriberRepository(ctrl)
		sender := mock_service.NewMockSender(ctrl)
		house := NewHouseService(hr, tm)
		subscriber := NewSubscriberService(sr, tm, sender)
		flat := NewFlatService(fr, house, subscriber, tm)

		fr.EXPECT().UpdateFlatStatus(gomock.Any(), req).Return(&ent2, nil)
		fr.EXPECT().GetFlatById(gomock.Any(), req.Id).Return(&ent1, nil)
		res, err := flat.UpdateFlat(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, mapper.FlatEntityToFlatResponse(&ent2), res)
	})
}
