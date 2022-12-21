import React from "react"
import Text from "../../models/text"
import "../preview/preview.css"

const TextViewer: React.FC<{ object: Text; [x: string]: any }> = ({
	object,
}) => {
	return (
		<div className="TextViewer">
			<div className="textHolder">
				<> {object.Value}</>
			</div>
		</div>
	)
}

export default TextViewer
