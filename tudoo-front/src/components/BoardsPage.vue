<template>
	<div class="boards-page">
		<h1>{{ msg }}</h1>
		<div class="board-section">
			<div>
				<ul class="board-section__list">
					<li class="board-section__list--item" v-for="(board, boardIndex) in boardList" :key="boardIndex">
						<div class="created-board" @mouseover="onMouseOver(boardIndex)" @mouseout="onMouseOut(boardIndex)">
							<div class="created-board__text">{{ board.name }}</div>
							<div class="created-board__link" :class="(boardLinkIndex === boardIndex) ? 'created-board__link--show' : 'created-board__link--hide'"><a :href="`/${board.name}`">Go</a></div>
						</div>
					</li>
					<li class="board-section__list--item">
						<div class="create-board" @click="showBoardCreator" v-if="!boardCreator">
							Create new board
						</div>
						<div class="add-board" v-if="boardCreator">
							<input type="text" v-model="boardTitle" class="create-board__input" maxlength="40" placeholder="Add board title">
							<button :disabled="!boardTitle" class="create-board__button" @click="createBoard">Create board</button>
						</div>
					</li>
				</ul>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: "BoardsPage",
	props: {
		msg: String
	},
	data () {
		return {
			boardCreator: false,
			boardTitle: '',
			boardList: [],
			boardLinkIndex: null
		}
	},
	methods: {
		showBoardCreator () {
			this.boardCreator = true
		},
		resetBoardTitle () {
			this.boardTitle = ''
		},
		createBoard () {
			let name = this.boardTitle
			let board = {}
			board.name = name
			this.boardList.push(board)
			this.resetBoardTitle()
		},
		onMouseOver (index) {
			this.boardLinkIndex = index
		},
		onMouseOut () {
			this.boardLinkIndex = null
		}
	}
}
</script>
