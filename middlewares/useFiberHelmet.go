package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
)

func UseFiberHelmet(app *fiber.App) {
	app.Use(helmet.New(helmet.Config{
		Filter:                    nil,                   // กำหนดฟังก์ชันเพื่อข้าม middleware ถ้าจำเป็น
		XSSProtection:             "1; mode=block",       // เปิดใช้งาน XSS Protection
		ContentTypeNosniff:        "nosniff",             // ป้องกันการ MIME-sniffing
		XFrameOptions:             "DENY",                // ป้องกันการแสดงผลใน iframe
		HSTSMaxAge:                31536000,              // เปิด HSTS เป็นเวลา 1 ปี
		HSTSExcludeSubdomains:     false,                 // รวม subdomains ใน HSTS หรือไม่
		ContentSecurityPolicy:     "default-src 'self';", // กำหนด CSP
		CSPReportOnly:             false,                 // ใช้ CSP ในโหมด report only หรือไม่
		HSTSPreloadEnabled:        true,                  // เปิดใช้งาน HSTS Preload
		ReferrerPolicy:            "no-referrer",         // กำหนดนโยบาย Referrer
		PermissionPolicy:          "geolocation=(self)",  // กำหนด Permissions Policy
		CrossOriginEmbedderPolicy: "require-corp",        // กำหนด Cross-Origin-Embedder-Policy
		CrossOriginOpenerPolicy:   "same-origin",         // กำหนด Cross-Origin-Opener-Policy
		CrossOriginResourcePolicy: "same-origin",         // กำหนด Cross-Origin-Resource-Policy
		OriginAgentCluster:        "?1",                  // กำหนด Origin-Agent-Cluster
		XDNSPrefetchControl:       "off",                 // กำหนด X-DNS-Prefetch-Control
		XDownloadOptions:          "noopen",              // ป้องกันการดาวน์โหลดไฟล์โดยไม่ตั้งใจ
		XPermittedCrossDomain:     "none",                // กำหนด X-Permitted-Cross-Domain-Policies
	}))
}
