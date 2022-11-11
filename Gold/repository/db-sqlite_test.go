package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {

	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate 失败: ", err)
	}

}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("插入失败: ", err)
	}

	if result.ID <= 0 {
		t.Error("交易id必须大于0: ", result.ID)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	rows, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("查询失败")
	}

	if len(rows) != 1 {
		t.Error("返回行数不为1，数量为: ", len(rows))
	}
}

func TestSQLiteRepository_GetHoldingByID(t *testing.T) {
	row, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("查询id失败: ", err)
	}
	if row.PurchasePrice != 1000 {
		t.Error("PurchasePrice的值不为1000: ", row.PurchasePrice)
	}
	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("无此id，查询错误:")
	}
}

func TestSQLiteRepository_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}
	h.PurchasePrice = 10001
	err = testRepo.UpdateHolding(1, *h)
	if err != nil {
		t.Error("更新失败: ", err)
	}
}

func TestSQLiteRepository_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("删除失败", err)

		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}
	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error when trying to delete non-existent record")
	}
}
