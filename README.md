## Vì sao bạn nên tự đặt tên folder hợp lý?

- ✅ Go import path phụ thuộc vào tên folder
- ✅ Giúp tổ chức code rõ ràng, dễ bảo trì
- ✅ Tự động mapping sang package name
- ✅ Tránh lỗi hoặc xung đột khi import

## Cấu trúc project

| Folder      | Nhiệm vụ                               | Package name       |
| ----------- | -------------------------------------- | ------------------ |
| `models/`   | Định nghĩa struct (User, Product...)   | `package models`   |
| `handlers/` | Xử lý request Fiber (`c *fiber.Ctx`)   | `package handlers` |
| `services/` | Logic xử lý nghiệp vụ (business logic) | `package services` |
| `routes/`   | Đăng ký route cho Fiber                | `package routes`   |
| `database/` | Kết nối, khởi tạo database             | `package database` |
| `config/`   | Đọc `.env`, biến cấu hình              | `package config`   |
| `storage/`  | Xử lý lưu local (JSON, in-memory...)   | `package storage`  |

## truyền by value trong go là gì ?

- Khi bạn truyền một biến vào hàm, Go sẽ sao chép (copy) giá trị của biến đó, chứ không truyền biến gốc vào.

## LƯU Ý

- Trong Go, chỉ có con trỏ (&) mới cho phép một hàm thay đổi giá trị thực sự của biến bên ngoài
