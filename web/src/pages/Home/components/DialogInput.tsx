interface DialogInputProps {
	label: string
	inputValue: string
	name: string
	onChange: (inputName: string, inputValue: string) => void
}
export function DialogInput({ label, inputValue, onChange, name }: DialogInputProps) {
	return (
		<div className="grid grid-cols-4 items-center gap-4">
			<label className="font-bold w-[100px] text-right">{label}</label>
			<input
				name={name}
				type="text"
				className="border border-zinc-600 rounded-sm outline-0 py-1 px-2 col-span-3 focus:outline-zinc-300 focus:outline-2"
				value={inputValue}
				onChange={(e) => onChange(name, e.target.value)}
			/>
		</div>
	)
}