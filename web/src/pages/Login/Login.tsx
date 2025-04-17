import { useState } from "react";
import { Button } from "../../components/Button";
import { Input } from "../../components/Input";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";
import { useUser } from "@/context/Context";

interface LoginFormData {
	email: string
	password: string
}

export function Login() {
	const [loginFormData, setLoginFormData] = useState<LoginFormData>({
		email: "",
		password: "",
	})
	const { setUserId } = useUser()
	const navigate = useNavigate()

	function handleLoginRequest(e: React.FormEvent) {
		e.preventDefault()
		const backendURL = import.meta.env.VITE_BACKEND_URL

		axios.post(`${backendURL}/login`, loginFormData, {
			withCredentials: true
		})
			.then((response) => {
				toast.success("Login successful", {
					style: {
						background: 'green',
						color: 'white',
						border: 'none',
						fontSize: 14,
					},
					duration: 3000
				})

				const userId = response.data.id
				setUserId(userId)

				navigate("/home")
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
		setLoginFormData(state => ({ ...state, [name]: value }))
	}

	return (
		<div className="flex min-h-screen">
			<div className="w-4xl bg-zinc-900 flex justify-center items-center py-16">
				<div className="flex flex-col gap-8 w-xl">
					<h1 className="text-4xl font-bold text-white">
						Login
					</h1>

					<form onSubmit={handleLoginRequest} className="flex flex-col gap-12 items-center">
						<div className="flex w-full flex-col gap-8 items-center">
							<Input
								inputType="email"
								label="Email address"
								name="email"
								onChange={handleInputChange}
								value={loginFormData.email}
								placeholder="example@example.com"
							/>
							<Input
								inputType="password"
								label="Password"
								name="password"
								onChange={handleInputChange}
								value={loginFormData.password}
								placeholder="Type your password"
							/>
						</div>
						<Button btnType="submit" btnContent="Login" />
					</form>
				</div>
			</div>

			<div className="flex-1 bg-purple-700 flex items-center justify-center py-16">
				<div className="flex flex-col gap-16 items-center w-3xl">
					<h1 className="text-white text-6xl font-bold text-center">
						New Here?
					</h1>
					<h2 className="text-white text-3xl text-center font-bold'">
						If you don't have an account yet, create one by clicking the button below.
					</h2>

					<Button btnType="navigate" btnContent="Signup" navigateTo="signup" />
				</div>

			</div>
		</div>
	)
}