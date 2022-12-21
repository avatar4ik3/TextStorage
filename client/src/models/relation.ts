export default class Relation {
	private _TextId: Number
	public get TextId(): Number {
		return this._TextId
	}
	public set TextId(v: Number) {
		this._TextId = v
	}

	private _GroupId: Number
	public get GroupId(): Number {
		return this._GroupId
	}
	public set GroupId(v: Number) {
		this._GroupId = v
	}

	constructor(text_id: Number, group_id: Number) {
		this._TextId = text_id
		this._GroupId = group_id
	}
}
