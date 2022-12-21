import React, { Key } from "react"
import Group from "../../models/group"
import Relation from "../../models/relation"
import Text from "../../models/text"
import Preview from "./preview"
import { PropsWithChildren } from "react"
import { TransitionGroup, CSSTransition } from "react-transition-group"

const PreviewList: React.FC<PropsWithChildren<{
	values: Array<Text | Group>
}>> = ({ values }) => {
	console.log(values)
	console.log("typeof values : ", typeof values)
	console.log("is array", Array.isArray(values))
	return (
		<div className="previewList">
			<TransitionGroup>
				{values.map((x) => (
					<CSSTransition key={x.id} classNames="previews" timeout={500}>
						<Preview object={x}></Preview>
					</CSSTransition>
				))}
			</TransitionGroup>
		</div>
	)
}

export default PreviewList
