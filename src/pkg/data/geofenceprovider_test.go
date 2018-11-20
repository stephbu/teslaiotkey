package data

import (
	"context"
	"github.com/stephbu/teslaiotkey/src/pkg/logging"
	"testing"
)

// Positive Test
func TestGetCarToFenceDistanceDrivewayPass(t *testing.T) {

	car := NewMockCarProvider(LatLong{Lat: 47.642744, Long: -122.112747}, false)
	fence := NewMockFenceProvider(LatLong{Lat: 47.642744, Long: -122.112782}, 30, false)

	ctx := logging.CreateRequestContext(context.Background())
	distance, err := FenceToPointDistance(ctx, fence, car)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if distance > fence.GetDistance() {
		t.Errorf("Distance incorrect: %v, want: <%v.", distance, fence.GetDistance())
	}
	t.Logf("Distance correct: %v", distance)
}

// Positive Test
func TestGetCarToFenceDistanceNearPass(t *testing.T) {

	car := NewMockCarProvider(LatLong{Lat: 47.642905, Long: -122.112694}, false)
	fence := NewMockFenceProvider(LatLong{Lat: 47.642759, Long: -122.112789}, 30, false)

	ctx := logging.CreateRequestContext(context.Background())
	distance, err := FenceToPointDistance(ctx, fence, car)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if distance > fence.GetDistance() {
		t.Errorf("Distance incorrect: %v, want: <%v.", distance, fence.GetDistance())
	}
	t.Logf("Distance correct: %v", distance)
}

// Negative test
func TestGetCarToFenceDistanceNeighbourhoodFail(t *testing.T) {

	car := NewMockCarProvider(LatLong{Lat: 47.643234, Long: -122.112018}, false)
	fence := NewMockFenceProvider(LatLong{Lat: 47.642759, Long: -122.112789}, 30, false)

	ctx := logging.CreateRequestContext(context.Background())
	distance, err := FenceToPointDistance(ctx, fence, car)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if distance < fence.GetDistance() {
		t.Errorf("neighbourhood distance should be ~78m, got %v", distance)
	}
	t.Logf("Distance correct: %v", distance)
}
