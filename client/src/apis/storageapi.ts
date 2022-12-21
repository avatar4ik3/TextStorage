import { Axios, AxiosResponse } from "axios"
import Group from "../models/group"
import Relation from "../models/relation"
import Text from "../models/text"

class DefaultCrud<T> {
	private endpoint: string

	private api: Axios

	constructor(endpoint: string, baseUrl: string) {
		this.endpoint = endpoint
		this.api = new Axios({
			baseURL: baseUrl,
			timeout: 1000,
			headers: {
				"Content-type": "application/json; charset=UTF-8",
			},
		})
	}

	async create(value: T): Promise<AxiosResponse<T, T>> {
		return this.api.post(this.endpoint, {
			params: value,
		})
	}

	async get(id: Number): Promise<AxiosResponse<T, T>> {
		return this.api.get(this.endpoint, {
			params: {
				id: id,
			},
		})
	}

	async getAll(): Promise<AxiosResponse<Array<T>, T>> {
		return this.api.get(this.endpoint)
	}

	async delete(id: Number): Promise<AxiosResponse<any, any>> {
		return this.api.delete(this.endpoint, {
			data: {
				id: id,
			},
		})
	}

	async deleteAll() {
		return this.api.delete(this.endpoint)
	}
}

export class TextApi extends DefaultCrud<Text> {
	constructor(baseUrl: string) {
		super("/texts", baseUrl)
	}
}

export class GroupsApi extends DefaultCrud<Group> {
	constructor(baseUrl: string) {
		super("/texts", baseUrl)
	}
}

export class RelationsApi extends DefaultCrud<Relation> {
	constructor(baseUrl: string) {
		super("/texts", baseUrl)
	}
}
