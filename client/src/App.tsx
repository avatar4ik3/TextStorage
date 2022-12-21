import React, { PropsWithChildren, useEffect, useState } from "react"
import { TextApi } from "./apis/storageapi"
import PreviewList from "./components/preview/previewList"
import Group from "./models/group"
import Text from "./models/text"

const App: React.FC<PropsWithChildren<{ api: TextApi }>> = ({ api }) => {
	const [texts, setTexts] = useState(Array<Text | Group>(5))
	const [isTextLoaded, setIsTextLoaded] = useState(false)
	useEffect(() => {
		setIsTextLoaded(false)
		api
			.getAll()
			.then((res) => setTexts(JSON.parse((res.data as unknown) as string)))
			.then(() => setIsTextLoaded(true))

		return () => {}
	}, [])

	return (
		<div className="App">
			{isTextLoaded ? <PreviewList values={texts}></PreviewList> : "getting"}
		</div>
	)
}

export default App
