import { BrowserRouter } from "react-router-dom";
import { Router } from "./Router";
import { Toaster } from "sonner";
import { UserProvider } from "./context/UserContext";

export function App() {

	return (
		<BrowserRouter>
			<Toaster position="top-right" />
			<UserProvider>
				<Router />
			</UserProvider>
		</BrowserRouter>
	)
}

