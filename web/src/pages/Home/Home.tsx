import { Header } from "../../components/Header";
import { UserData } from "./components/UserData";
import { useEffect, useState } from "react";
import axios from "axios";
import { toast } from "sonner";
import { CloseButton, Dialog, Portal, Text } from "@chakra-ui/react"
import { DialogForm } from "./components/DialogForm";
import { useUser } from "../../context/Context";

export interface UserResponse {
	Name: string,
	Email: string,
	Cpf: string,
	Age: string,
	PhoneNumber: string,
}

export function Home() {
	const { userId } = useUser()
	const [userResponse, setUserResponse] = useState<UserResponse>({
		Age: '',
		Cpf: '',
		Name: '',
		Email: '',
		PhoneNumber: ''
	})
	const [isDialogOpen, setIsDialogOpen] = useState(false)

	useEffect(() => {
		if (!userId) return

		function fetchUserData() {
			const backendURL = import.meta.env.VITE_BACKEND_URL
			axios.get<UserResponse>(`${backendURL}/user/${userId}`, {
				withCredentials: true
			})
				.then((response) => {
					response.data.Cpf = response.data.Cpf.replace(/^(\d{3})(\d{3})(\d{3})(\d{2})$/, "$1.$2.$3-$4")
					setUserResponse(response.data)
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

		fetchUserData()
	}, [userId])

	return (
		<div className="h-screen flex flex-col">
			<Header />
			<div className="flex flex-col flex-1 items-center justify-center text-white bg-zinc-900">
				<h1 className="text-4xl font-bold mb-10 text-center">
					Welcome back, <span className="text-purple-500">{userResponse.Name}</span>
				</h1>

				<div className="bg-zinc-800 flex flex-col gap-2 items-center rounded-2xl shadow-lg p-8 w-full max-w-xl">
					<div className="w-full flex flex-col gap-2">
						<UserData label="Email" data={userResponse.Email} />
						<UserData label="CPF" data={userResponse.Cpf} />
						<UserData label="Age" data={userResponse.Age} />
						<UserData label="Phone Number" data={userResponse.PhoneNumber} />
					</div>

					<Dialog.Root placement="center" lazyMount open={isDialogOpen} onOpenChange={(e) => setIsDialogOpen(e.open)}>
						<Dialog.Trigger asChild>
							<button className="px-6 py-3 bg-purple-500 rounded-sm hover:bg-purple-400 transition delay-100 font-bold cursor-pointer">
								Editar Perfil
							</button>
						</Dialog.Trigger>
						<Portal>
							<Dialog.Backdrop />
							<Dialog.Positioner>
								<Dialog.Content bg="gray.900" color="white" borderColor="gray.700" p={6} maxW="lg">
									<Dialog.Header>
										<Dialog.Title fontSize="xl" fontWeight="bold">
											Edit profile
										</Dialog.Title>
									</Dialog.Header>

									<Dialog.Body>
										<Text color="gray.400" mb={12}>
											Make changes to your profile here. Click save when you're done.
										</Text>

										<DialogForm
											name={userResponse.Name}
											CPF={userResponse.Cpf}
											email={userResponse.Email}
											age={userResponse.Age}
											phoneNumber={userResponse.PhoneNumber}
											setUserResponse={setUserResponse}
											setIsDialogOpen={setIsDialogOpen}
										/>
									</Dialog.Body>

									<Dialog.CloseTrigger asChild>
										<CloseButton size="sm" />
									</Dialog.CloseTrigger>
								</Dialog.Content>
							</Dialog.Positioner>
						</Portal>
					</Dialog.Root>

				</div>
			</div>
		</div>
	)
}