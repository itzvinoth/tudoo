<template>
	<div class="tasks">
		<h1>This is an tasks page</h1>
		<draggable
			:list="contents"
			:disabled="!enabled"
			item-key="value"
			class="list-group"
			ghost-class="ghost"
			:move="checkMove"
			@start="dragging = true"
			@end="dragging = false"
		>
			<template #item="{ element }">
				<div class="list-group-item" :class="{ 'not-draggable': !enabled }">
					<div
						:id="`content-${element.index}`"
						class="content"
						contenteditable
						spellcheck="false"
						@input="event => onInput(event, element.index)"
						@keyup.delete="onRemove(element.index)"
						@keydown.enter="onEnter"
						@keydown.enter.prevent
					/>
				</div>
			</template>
		</draggable>
	</div>
</template>

<script>
import draggable from "vuedraggable"

export default {
	name: "Tasks",
	components: {
		draggable
	},
	data () {
		return {
			enabled: true,
			contents: [
				{ value: 'Next adding drag and drop', index: 0 },
				{ value: 'Alignment needs to be fixed', index: 1 },
				{ value: 'CSS has to be updated', index: 2 },
				{ value: '', index: 3 },
			]
		}
	},
	// created () {
	// 	const plugin = document.createElement('script')
	// 	plugin.setAttribute(
	// 		'src',
	// 		'../assets/js/dnd.js'
	// 	)
	// 	plugin.async = true
	// 	document.head.appendChild(plugin)
	// },
	mounted () {
		this.updateContents()
	},
	methods: {
		onInput (event, index) {
			const value = event.target.innerText
			this.contents[index].value = value
			this.contents[index].index = this.contents.length - 1
			console.log('contents: ', this.contents)
		},
		onEnter (event) {
			console.log('index: ', event.target.innerText)
			if (event.target.innerText) {
				this.contents.push({value: '', index: this.contents.length - 1})
				this.$nextTick(() => {
					let totalContents = this.contents.length - 1
					let el = document.querySelector('#content-'+(totalContents))
					el.focus()
					// event.target.nextElementSibling.focus()
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

<style scoped>
.buttons {
  margin-top: 35px;
}
.ghost {
  opacity: 0.5;
  background: #c8ebfb;
}
.not-draggable {
  cursor: no-drop;
}
</style>