import { TruncateText } from "@/utils/string"

interface UserInfoProps {
	label: string
	content: string
}

export function UserInfo({ label, content }: UserInfoProps) {
	return (
		<div className="flex gap-2 ">
			<h1 className="text-xl text-zinc-400">{label}:</h1>
			<h1 className="text-lg text-zinc-200" title={content}>{TruncateText(content)}</h1>
		</div>
	)

}