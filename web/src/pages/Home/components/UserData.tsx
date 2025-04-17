interface UserDataProps {
	label: string
	data: string
}

export function UserData({ label, data }: UserDataProps) {
	return (
		<div className="flex flex-col">
			<p className="text-sm text-zinc-400 uppercase">{label}:</p>
			<p>{data}</p>
		</div>
	)
}