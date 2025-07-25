package bootcampsrvc

import (
   _ "log"
	"context"
	_"database/sql"
    // _"net/http"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
    "github.com/melikadaryabak/bootcamp/internal/infrastructure/db"
)

type BootcampSrvc struct {
    repo *db.BootcampRepo
}

func NewBootcampSrvc(repo *db.BootcampRepo) *BootcampSrvc {
	return &BootcampSrvc{repo: repo}
}

func (b BootcampSrvc) GetBootcamps(ctx context.Context) ([]entity.Bootcamp, error) {
	return b.repo.GetBootcamps(ctx)
}

func (b BootcampSrvc) PostBootcamp(ctx context.Context, bootcamp entity.Bootcamp) (int64, error) {
    return b.repo.PostBootcamps(ctx, bootcamp)
}

func (b BootcampSrvc) DeleteBootcamp(ctx context.Context,id int64) (bool, error) {
    return b.repo.DeleteBootcamps(ctx, id)
}

func (b BootcampSrvc) PutBootcamp(ctx context.Context,bootcamp entity.Bootcamp) (bool, error) {
    return b.repo.PutBootcamps(ctx, bootcamp)
}