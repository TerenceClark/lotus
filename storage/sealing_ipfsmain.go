package storage

import (
	"context"
	"github.com/filecoin-project/specs-actors/actors/abi"
	sealing "github.com/filecoin-project/storage-fsm"
	"io"
)

func (m *Miner) AllocatePieceAndSendIfNeeded(size abi.UnpaddedPieceSize) (sectorID abi.SectorNumber, offset uint64, err error) {
	return m.sealing.AllocatePieceAndSendIfNeeded(size)
}

func (m *Miner) AddPiece(ctx context.Context, size abi.UnpaddedPieceSize, r io.Reader, sectorID abi.SectorNumber, d sealing.DealInfo) error {
	return m.sealing.AddPiece(ctx, size, r, sectorID, d)
}
