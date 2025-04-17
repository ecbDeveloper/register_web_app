import { useNavigate } from 'react-router-dom'

type btnType = 'submit' | 'navigate'
type navigateToType = 'login' | 'signup'

interface ButtonProps {
	btnType: btnType
	navigateTo?: navigateToType
	btnContent: string
}


export function Button({ btnType, btnContent, navigateTo }: ButtonProps) {
	const navigate = useNavigate()

	const navigatePath = navigateTo === 'login' ? '/' : '/signup'

	return (
		btnType === 'submit' ? (
			<button type='submit' className='w-96 py-3 bg-purple-700 rounded-full border border-purple-500 text-white text-3xl cursor-pointer hover:bg-purple-500'>
				{btnContent}
			</button>
		) : (
			<button type='button' onClick={() => navigate(navigatePath)} className='w-96 py-3 bg-zinc-900 rounded-full border border-zinc-600 text-white text-3xl cursor-pointer hover:bg-zinc-700'>
				{btnContent}
			</button>
		)
	)
}