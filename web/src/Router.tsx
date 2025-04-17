import { Route, Routes } from "react-router-dom";
import { Signup } from "./pages/Signup/Signup";
import { Login } from "./pages/Login/Login";
import { Home } from "./pages/Home/Home";
import { PrivateRoute } from "./routes/PrivateRoute";
import { PublicOnlyRoute } from "./routes/PublicOnlyRoute";
import { Dashboard } from "./pages/Dashboard/Dashboard";

export function Router() {
	return (
		<Routes>

			<Route path='/' element={
				<PublicOnlyRoute>
					<Login />
				</PublicOnlyRoute>

			} />
			<Route path='/signup' element={
				<PublicOnlyRoute>
					<Signup />
				</PublicOnlyRoute>
			} />

			<Route path='/home' element={
				<PrivateRoute>
					<Home />
				</PrivateRoute>
			} />

			<Route path="/admin/dashboard" element={
				<PrivateRoute>
					<Dashboard />
				</PrivateRoute>
			} />

		</Routes>
	)
}