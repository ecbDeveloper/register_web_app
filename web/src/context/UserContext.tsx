import { useState, useEffect, PropsWithChildren } from "react";
import { UserContext } from "./Context";
import axios from "axios";
import { getItemWithExpiry, setItemWithExpiry } from "../utils/storage";


export const UserProvider = ({ children }: PropsWithChildren) => {
	const [userId, setUserIdState] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const [isAdmin, setIsAdmin] = useState<boolean | null>(null)


	useEffect(() => {
		const backendURL = import.meta.env.VITE_BACKEND_URL

		axios.get(`${backendURL}/auth/admin`, { withCredentials: true })
			.then(() => setIsAdmin(true))
			.catch(() => setIsAdmin(false));

		const storedId = getItemWithExpiry();
		setUserIdState(storedId);
		setIsLoading(false)
	}, []);

	const setUserId = (id: string | null) => {
		if (id) {
			setItemWithExpiry(id);
			setUserIdState(id);
		} else {
			localStorage.removeItem("userId");
			setUserIdState(null);
		}
	};

	return (
		<UserContext.Provider value={{ userId, setUserId, isLoading, isAdmin }}>
			{children}
		</UserContext.Provider>
	);
};