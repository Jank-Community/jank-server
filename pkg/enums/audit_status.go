package enums

type AuditStatus string

// AuditStatus 审核状态枚举
const (
	AuditPending  AuditStatus = "pending"  // 待审核
	AuditApproved AuditStatus = "approved" // 审核通过
	AuditRejected AuditStatus = "rejected" // 审核不通过
)
