<script>
	import UserIcon from '../../assets/icons/User.png'
	import TrelloIcon from '../../assets/icons/trello-icon.png'
	import LogoTrello from '../../assets/icons/logo-trello.png'
	import PlusCircleOutline from '../../assets/icons/plus-circle-outline.png'
	import AlertCircleOutline from '../../assets/icons/alert-circle-outline.png'
	import BellOutline from '../../assets/icons/bell-outline.png'
	import SearchOutline from '../../assets/icons/search-outline.png'
</script>

<nav>
	<div class='row p-0 m-0'>
		<div
			class='d-flex justify-content-start align-items-center'
			style='height: 60px; border-bottom:1px solid #dcdcdc;'
		>
			<div class='col-lg-6 d-flex align-items-center'>
				<a href='/'>
					<img alt='logo trello' src={LogoTrello}/>
				</a>
				<span class='separator'>|</span>
				<a href='/boards' style='text-decoration: none; color:black;'>
					<div class='d-flex'>
						<img alt='logo icon' src={TrelloIcon}/>
						<span>Boards</span>
					</div>
				</a>
				<span class='separator'>|</span>
				<div
					class='w100 d-flex justify-content-center align-items-center'
					style='position: relative;'
				>
					<input class='search-input' type='text'/>
					<img alt='search' class='search-icon' src={SearchOutline}/>
				</div>
			</div>
			<div class='col-lg-6 d-flex justify-content-end align-items-center'>
				<div
					class='d-flex justify-content-center align-items-center icons-right'
				>
					<div class='more'>
						<img alt='plus' src={PlusCircleOutline}/>
					</div>
					<div class='attention'>
						<img alt='allert' src={AlertCircleOutline}/>
					</div>
					<div class='bell'>
						<img alt='bell' src={BellOutline}/>
					</div>
					<div class='user'>
						<img alt='user' src={UserIcon}/>
					</div>
				</div>
			</div>
		</div>
	</div>
</nav>

<style>
	.w100 {
		width: 100%;
	}

	.separator {
		color: #dcdcdc;
		margin: 0 30px;
	}

	.search-input {
		background-color: #dcdcdc;
		border: none;
		border-radius: 25px;
		width: 100%;
		padding: 0 50px;
		height: 30px;
	}

	.search-icon {
		position: absolute;
		right: 0;
		margin-right: 30px;
	}

	.icons-right > div {
		margin: 0 15px;
		cursor: pointer;
	}
</style>
