package mysql

import (
	"fmt"
	"testing"
)

func TestGetRecordList(t *testing.T) {
	// 调用测试函数
	page := 1
	size := 10
	userID := 1
	records, err := GetRecordList(int64(page), int64(size), int64(userID))
	if err != nil {
		t.Fatalf("failed to get records: %v", err)
	}
	fmt.Println("Records:")
	for _, record := range records {
		fmt.Printf("ID: %d, DiseaseName: %s, Alias: %s, CreateTime: %s\n", record.ID, record.DiseaseName, record.Alias, record.CreateTime)
	}
	// 检查结果是否符合预期
	expectedRecordCount := 1
	if len(records) != expectedRecordCount {
		t.Fatalf("unexpected record count: got %d, want %d", len(records), expectedRecordCount)
	}
}
