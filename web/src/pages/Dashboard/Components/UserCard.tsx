import { UserResponse } from "../../Home/Home";
import { UserInfo } from "./UserInfo";

export function UserCard({ Age, Cpf, Email, Name, PhoneNumber }: UserResponse) {
	return (
		<div className="bg-zinc-800 w-fit h-50 flex flex-col gap-2 justify-center p-8 rounded-lg">
			<UserInfo label="Name" content={Name} />
			<UserInfo label="Email" content={Email} />
			<UserInfo label="Age" content={Age} />
			<UserInfo label="CPF" content={Cpf} />
			<UserInfo label="Phone number" content={PhoneNumber} />
		</div>
	)
}