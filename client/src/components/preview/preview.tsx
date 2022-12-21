import React from "react"
import Group from "../../models/group"
import Text from "../../models/text"
import { PropsWithChildren } from "react"
import "../preview/preview.css"

const Preview: React.FC<PropsWithChildren<{ object: Text | Group }>> = ({
	object,
}) => {
	return (
		<div className="preview">
			<div className="name">
				<> name : {object.description}</>
			</div>
			<div className="description">
				<> id : {object.id} </>
			</div>
		</div>
	)
}

export default Preview
