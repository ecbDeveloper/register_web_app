import { useEffect, useState } from "react";
import { UserResponse } from "../Home/Home";
import axios from "axios";
import { toast } from "sonner";
import { useNavigate } from "react-router-dom";
import { UserCard } from "./Components/UserCard";
import { useUser } from "../../context/Context";
import { Header } from "../../components/Header";

export function Dashboard() {
	const { isAdmin, isLoading } = useUser()
	const navigate = useNavigate()
	const [usersResponse, setUsersResponse] = useState<UserResponse[]>([])

	useEffect(() => {
		if (isLoading) return

		if (isAdmin) {
			function fetchUsersData() {
				const backendURL = import.meta.env.VITE_BACKEND_URL
				axios.get<UserResponse[]>(`${backendURL}/admin/users`, {
					withCredentials: true
				})
					.then((response) => {
						const responseData = response.data.map((obj) => ({
							...obj,
							Cpf: obj.Cpf.replace(/^(\d{3})(\d{3})(\d{3})(\d{2})$/, "$1.$2.$3-$4")
						}))

						setUsersResponse(responseData)
					})
					.catch((err) => {
						toast.error(`Failed to fetch user data: ${JSON.stringify(err.response.data.message)}`, {
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

			fetchUsersData()
		}
	}, [isAdmin, navigate, isLoading])

	return (
		<div className="h-screen flex flex-col bg-zinc-900 ">
			<Header />
			<div className="mt-20 flex flex-col gap-20 items-center">
				<h1 className="text-5xl text-white font-bold w-fit">Usuarios</h1>
				<div className="w-full grid grid-cols-4 justify-items-center text-white">
					{usersResponse.map((obj, index) => (
						<UserCard key={index} Age={obj.Age} Name={obj.Name} Cpf={obj.Cpf} Email={obj.Email} PhoneNumber={obj.PhoneNumber} />
					))}
				</div>
			</div>
		</div>
	)
}