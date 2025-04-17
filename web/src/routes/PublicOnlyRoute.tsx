import { Navigate } from "react-router-dom";
import { useUser } from "@/context/Context";

export function PublicOnlyRoute({ children }: { children: React.ReactNode }) {
	const { userId, isLoading } = useUser()

	if (isLoading) {
		return (
			<div className="h-screen bg-zinc-900 flex justify-center items-center">
				<h1 className="text-3xl text-white">Carregando...</h1>
			</div>
		)
	}

	return userId ? <Navigate to="/home" /> : children;
}