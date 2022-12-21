import React, { Key } from "react"
import Group from "../../models/group"
import Relation from "../../models/relation"
import Text from "../../models/text"
import Preview from "./preview"
import { PropsWithChildren } from "react"
import { TransitionGroup, CSSTransition } from "react-transition-group"

const PreviewList: React.FC<PropsWithChildren<{
	values: Array<Text>
	set: (value: Text) => void
}>> = ({ values, set }) => {
	set(values[0])
	return (
		<div className="previewList">
			{values.map((x) => {
				return (
					<Preview
						key={x.Id.toString()}
						object={x}
						handler={() => set(x)}
					></Preview>
				)
			})}
		</div>
	)
}

export default PreviewList
