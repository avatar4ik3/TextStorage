export default class Relation {
	private _textId: Number
	public get textId(): Number {
		return this._textId
	}
	public set textId(v: Number) {
		this._textId = v
	}

	private _groupId: Number
	public get groupId(): Number {
		return this._groupId
	}
	public set groupId(v: Number) {
		this._groupId = v
	}

	constructor(text_id: Number, group_id: Number) {
		this._textId = text_id
		this._groupId = group_id
	}
}
