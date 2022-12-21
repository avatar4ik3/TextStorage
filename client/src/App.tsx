import { time } from "console"
import React, { PropsWithChildren, useEffect, useState } from "react"
import { TextApi } from "./apis/storageapi"
import PreviewList from "./components/preview/previewList"
import Group from "./models/group"
import Text from "./models/text"

const App: React.FC<PropsWithChildren<{ api: TextApi }>> = ({ api }) => {
	const [texts, setTexts] = useState<Array<Text>>()
	const [selected, setSelected] = useState<Text | undefined>()
	console.log("type", typeof setSelected)
	const [isTextLoaded, setIsTextLoaded] = useState<boolean>(false)
	useEffect(() => {
		setIsTextLoaded(false)
		api
			.getAll()
			.then(async (res) => {
				setTexts(JSON.parse((res.data as unknown) as string))
				console.log(texts)
			})
			.then(() => setIsTextLoaded(true))

		return () => {}
	}, [])

	return (
		<div className="App">
			<>
				{isTextLoaded ? (
					<div>
						<PreviewList
							values={texts as Text[]}
							set={setSelected}
						></PreviewList>
						<div>current state</div>
						<div>
							{selected !== undefined ? (
								<>
									{selected?.Id} {selected?.Description}
								</>
							) : (
								"load"
							)}
						</div>
					</div>
				) : (
					"getting"
				)}
			</>
		</div>
	)
}

export default App
