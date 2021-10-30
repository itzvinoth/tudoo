<template>
	<div class="tasks">
		<h1>This is an tasks page</h1>
		<p
			v-for="(value, index) in contents"
			:id="`content-${index}`"
			class="content"
			:key="index"
			contenteditable
			spellcheck="false"
			@input="event => onInput(event, index)"
			@keyup.delete="onRemove(index)"
			@keydown.enter="onEnter"
			@keydown.enter.prevent
		/>
	</div>
</template>

<script>
export default {
	name: "Tasks",
	data () {
		return {
			contents: [
				{ value: 'paragraph 1' },
				{ value: 'paragraph 2' },
				{ value: 'paragraph 3' },
				{ value: '' },
			]
		}
	},
	mounted () {
		this.updateContents()
	},
	methods: {
		onInput (event, index) {
			const value = event.target.innerText
			this.contents[index].value = value
		},
		onEnter (event) {
			// console.log('event: ',)
			// console.log('index: ', event.target.innerText)
			if (event.target.innerText) {
				this.contents.push({value: ''})
				this.$nextTick(() => {
					event.target.nextElementSibling.focus()
				})
			} else {
				return			
			}
		},
		onRemove (index) {
			console.log(index)
			// if (this.contents.length > 1 && this.contents[index].value.length === 0) {
			// 	this.$delete(this.contents, index)
			// 	this.updateContents()
			// }
		},
		updateContents () {
			this.contents.forEach((content, index) => {
				const el = document.getElementById(`content-${index}`)
				el.innerText = content.value
			})
		}
	}
}
</script>
