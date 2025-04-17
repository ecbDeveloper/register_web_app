# 🚀 Register System

This is a user registration system with authentication and authorization, developed with a separate frontend and backend architecture. The project includes features such as user registration, login, profile editing, and user management for administrators.

## 🛠️ Technologies Used

### Backend
- **Language**: Go 🐹
- **Framework**: Echo ⚡
- **Database**: PostgreSQL 🐘
- **SQL Manager**: sqlc 📜
- **Authentication**: JWT (JSON Web Tokens) 🔑
- **Validation**: Ozzo Validation ✅
- **Documentation**: Swagger 📖

### Frontend
- **Language**: TypeScript 🟦
- **Framework**: React ⚛️
- **Routing**: React Router DOM 🛣️
- **Styling**: TailwindCSS 🎨
- **State Management**: Context API 🌐
- **Additional Libraries**:
  - Axios (HTTP requests) 📡
  - React Query (asynchronous data management) 🔄
  - Sonner (notifications) 🔔

## ✨ Features
- **User**:
  - Registration 📝
  - Login 🔓
  - Profile Editing ✏️
- **Administrator**:
  - User Listing 📋
  - User Data Management 🛠️

## 🏃 How to Run

### Backend
1. Ensure Docker is installed 🐳.
2. Start the database with:
   ```bash
   docker-compose up
   ```
3. Configure environment variables in the `.env` file 🛠️.
4. Run the backend:
   ```bash
   go run main.go
   ```

### Frontend
1. Install dependencies 📦:
   ```bash
   npm install
   ```
2. Start the development server 🚀:
   ```bash
   npm run dev
   ```

## 🤝 Contribution
Contributions are welcome! Feel free to open issues or submit pull requests. 💡