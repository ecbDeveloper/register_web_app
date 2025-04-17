import { createContext, useContext } from "react";

type UserContextType = {
	userId: string | null;
	setUserId: (id: string | null) => void;
	isAdmin: boolean | null;
	isLoading: boolean
};

export const UserContext = createContext({} as UserContextType);

export const useUser = () => useContext(UserContext);