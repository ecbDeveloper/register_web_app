import { useUser } from "@/context/Context";
import axios from "axios";
import { LogOut } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

export function Header() {
	const { setUserId } = useUser();
	const navigate = useNavigate()

	function handleLogout() {
		const backendURL = import.meta.env.VITE_BACKEND_URL

		axios.post(`${backendURL}/auth/logout`, null, {
			withCredentials: true,
		})
			.then(() => {
				setUserId(null)
				navigate("/")

				toast("User logout successfully!", {
					style: {
						border: 'none',
						fontSize: 14,
					},
					duration: 3000
				})
			})
			.catch(() => {
				toast("Failed to logout, try again!", {
					style: {
						background: 'red',
						color: 'white',
						border: 'none',
						fontSize: 14,
					},
					duration: 3000,
				})
			})
	}

	return (
		<header className="bg-zinc-900 h-20 flex items-center justify-between px-10 border-b-1 border-b-zinc-600 ">
			<button onClick={() => navigate('/home')} className="flex gap-1 text-2xl text-white font-bold cursor-pointer">
				<p>Web</p>
				<p>Register</p>
			</button>

			<button
				className="group flex gap-1 items-center cursor-pointer"
				title="Logout"
				onClick={handleLogout}
			>
				<LogOut size={32} className="text-white group-hover:text-purple-700 transition delay-100" />
				<p className="text-xl text-white group-hover:text-purple-700 font-bold transition delay-100">Logout</p>
			</button>
		</header>
	)
}