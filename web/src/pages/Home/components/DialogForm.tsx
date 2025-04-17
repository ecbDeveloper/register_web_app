import axios from "axios";
import { DialogInput } from "./DialogInput";
import { FormEvent, useState } from "react";
import { useUser } from "@/context/Context";
import { toast } from "sonner";
import { UserResponse } from "../Home";

interface DialogFormProps {
	name: string
	CPF: string
	age: string
	email: string
	phoneNumber: string
	setUserResponse: React.Dispatch<React.SetStateAction<UserResponse>>
	setIsDialogOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export function DialogForm({ name, CPF, age, email, phoneNumber, setUserResponse, setIsDialogOpen }: DialogFormProps) {
	const { userId } = useUser()
	const [dialogFormData, setDialogFormData] = useState<DialogFormProps>({
		name: name,
		CPF: CPF,
		age: age,
		email: email,
		phoneNumber: phoneNumber,
		setUserResponse,
		setIsDialogOpen
	})

	// para cada objeto em dailofFormData verificar se o valor desse objeto é diferente do da propriedade recebida.

	function HandleInputChange(inputName: string, inputValue: string) {
		setDialogFormData((state) => ({ ...state, [inputName]: inputValue }))
	}

	function UpdateUserData() {
		const backendURL = import.meta.env.VITE_BACKEND_URL
		axios(`${backendURL}/user/${userId}`, {
			method: "put",
			data: dialogFormData,
			withCredentials: true,
		})
			.then(() => {
				setUserResponse({
					Age: dialogFormData.age,
					Cpf: dialogFormData.CPF,
					Email: dialogFormData.email,
					Name: dialogFormData.name,
					PhoneNumber: dialogFormData.phoneNumber
				})

				setIsDialogOpen(false)

				toast.success("User data updated successfully", {
					style: {
						background: 'green',
						color: 'white',
						border: 'none',
						fontSize: 14,
					},
					duration: 3000
				})
			})
			.catch((err) => {
				toast.error(`Failed to update user data: ${JSON.stringify(err.response.data.message)}`, {
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

	function handleUpdateUserInfo(e: FormEvent) {
		e.preventDefault()

		UpdateUserData()
	}

	return (

		<form onSubmit={handleUpdateUserInfo} className="flex flex-col gap-4 items-end">
			<DialogInput name="name" label="Name" inputValue={dialogFormData.name} onChange={HandleInputChange} />
			<DialogInput name="email" label="Email" inputValue={dialogFormData.email} onChange={HandleInputChange} />
			<DialogInput name="CPF" label="CPF" inputValue={dialogFormData.CPF} onChange={HandleInputChange} />
			<DialogInput name="age" label="Age" inputValue={dialogFormData.age} onChange={HandleInputChange} />
			<DialogInput name="phoneNumber" label=" Phone number" inputValue={dialogFormData.phoneNumber} onChange={HandleInputChange} />
			<button className="px-5 py-2.5 bg-white text-zinc-900 rounded-sm hover:bg-zinc-300 transition delay-100 cursor-pointer">
				Save changes
			</button>
		</form>

	)
}