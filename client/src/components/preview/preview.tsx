import React from "react"
import Group from "../../models/group"
import Text from "../../models/text"
import { PropsWithChildren } from "react"
import "../preview/preview.css"

const Preview: React.FC<{
	object: Text
	handler: React.MouseEventHandler<HTMLDivElement>
	[x: string]: any
}> = ({ object, handler }) => {
	return (
		<div className="preview" onClick={handler}>
			<div className="name">
				<> name : {object.Description}</>
			</div>
			<div className="id">
				<> id : {object.Id} </>
			</div>
		</div>
	)
}

export default Preview
