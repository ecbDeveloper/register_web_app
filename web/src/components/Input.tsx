import { Eye, EyeClosed } from "lucide-react"
import { JSX, useState } from "react"
import { IMaskInput } from "react-imask"

type inputType = "email" | "password" | "text"

interface InputProps {
	name: string
	label: string
	inputType: inputType
	onChange: (name: string, value: string) => void
	value: string
	placeholder: string
	mask?: string
}

export function Input({ label, inputType, name, onChange, value, placeholder, mask }: InputProps) {
	const [showPassword, setShowPassword] = useState(false)
	const [isClicked, setIsClicked] = useState(false)

	let icon: JSX.Element

	switch (inputType) {
		case "password": {
			const Icon = showPassword ? Eye : EyeClosed;

			icon =
				<button
					type="button"
					className="absolute right-4 top-2.5 cursor-pointer"
					title={showPassword ? "Hide password" : "Show password"}
					onClick={() => setShowPassword(!showPassword)}
				>
					<Icon size={28} className={isClicked ? "text-purple-500" : "text-white"} />
				</button>

			break
		}
		default:
			icon = <></>
	}

	const actualInputType = inputType === "password" && showPassword ? "text" : inputType

	return (
		<div className="flex flex-col gap-2 h-24 w-full">
			<label className="text-white text-xl">{label}</label>
			<div className="relative h-12">

				{mask ? (
					<IMaskInput
						mask={mask}
						value={value}
						onAccept={(val: string) => onChange(name, val)}
						name={name}
						type={actualInputType}
						placeholder={placeholder}
						onFocus={() => setIsClicked(true)}
						onBlur={() => setIsClicked(false)}
						className="focus:outline-0 focus:border-purple-700 focus:border rounded-xl px-6 py-4 relative bg-zinc-800 w-full text-white h-full text-xl"
						required
					/>
				) : (
					<input
						required
						name={name}
						value={value}
						type={actualInputType}
						placeholder={placeholder}
						onChange={(e) => onChange(name, e.target.value)}
						onFocus={() => setIsClicked(true)}
						onBlur={() => setIsClicked(false)}
						className="focus:outline-0 focus:border-purple-700 focus:border rounded-xl px-6 py-4 relative bg-zinc-800 w-full text-white h-full text-xl"
					/>
				)
				}

				{icon}
			</div>
		</div>
	)
}