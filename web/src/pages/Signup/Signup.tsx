import { useState } from "react";
import { Button } from "../../components/Button";
import { Input } from "../../components/Input";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

interface SignupFormData {
	name: string
	age: string
	cpf: string
	phoneNumber: string
	email: string
	password: string
}

export function Signup() {
	const [signupFormData, setSignupFormData] = useState<SignupFormData>({
		name: "",
		email: "",
		password: "",
		phoneNumber: "",
		cpf: "",
		age: ""
	})
	const navigate = useNavigate()

	function handleSignupRequest(e: React.FormEvent) {
		e.preventDefault()
		const backendURL = import.meta.env.VITE_BACKEND_URL

		const formattedDataToRequest = {
			...signupFormData,
			age: Number(signupFormData.age)
		}

		axios.post(`${backendURL}/signup`, formattedDataToRequest, {
			withCredentials: true
		})
			.then(() => {
				toast.success("Registered successfuly", {
					style: {
						background: 'green',
						color: 'white',
						border: 'none',
						fontSize: 14,
					},
					duration: 3000
				})
				navigate("/")
			})
			.catch((err) => {
				toast.error(`Login failed: ${JSON.stringify(err.response.data.message)}`, {
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

	function handleInputChange(name: string, value: string) {
		setSignupFormData(state => ({ ...state, [name]: value }))
	}

	return (
		<div className="flex min-h-screen">
			<div className="w-4xl bg-zinc-900 flex justify-center items-center py-16">
				<div className="flex flex-col gap-8 w-xl">
					<h1 className="text-4xl font-bold text-white">
						Signup
					</h1>

					<form onSubmit={handleSignupRequest} className="flex flex-col gap-8 items-center">
						<div className="flex w-full flex-col gap-8 items-center">
							<Input
								inputType="text"
								label="Name"
								name="name"
								value={signupFormData.name}
								onChange={handleInputChange}
								placeholder="Type your name"
							/>
							<Input
								inputType="text"
								label="Phone Number"
								name="phoneNumber"
								value={signupFormData.phoneNumber}
								onChange={handleInputChange}
								placeholder="( ) _____-____"
								mask="(00) 00000-0000"
							/>
							<Input
								inputType="text"
								label="Age"
								name="age"
								value={signupFormData.age}
								onChange={handleInputChange}
								placeholder="20"
							/>
							<Input
								inputType="text"
								label="CPF"
								name="cpf"
								value={signupFormData.cpf}
								onChange={handleInputChange}
								placeholder="000.000.000-00"
								mask="000.000.000-00"
							/>
							<Input
								inputType="email"
								label="Email address"
								name="email"
								value={signupFormData.email}
								onChange={handleInputChange}
								placeholder="example@example.com"
							/>
							<Input
								inputType="password"
								label="Password"
								name="password"
								value={signupFormData.password}
								onChange={handleInputChange}
								placeholder="Type your password"
							/>
						</div>

						<Button btnType="submit" btnContent="Signup" />
					</form>
				</div>
			</div>

			<div className="flex-1 bg-purple-700 flex justify-center py-80">
				<div className="flex flex-col gap-16 items-center w-3xl">
					<h1 className="text-white text-6xl font-bold text-center">
						Welcome back!
					</h1>
					<h2 className="text-white text-3xl text-center font-bold'">
						Already have an account? Log in below.
					</h2>

					<Button btnType="navigate" btnContent="Login" navigateTo="login" />
				</div>

			</div>
		</div>
	)
}