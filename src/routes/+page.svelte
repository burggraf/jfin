<script lang="ts">
    let currentIndex = -1;
    let currentItem: any;
	document.addEventListener('keydown', function (event) {
		// Get all focusable elements
		const focusableElements = document.querySelectorAll('.item')
        const focusableArray = Array.from(focusableElements); // Convert NodeList to Array

		// Find the index of the currently focused element
		let nextIndex = 0 // Initialize the nextIndex variable
        let e: KeyboardEvent;
		switch (event.keyCode) {
			case 37: // Left
			case 38: // Up
                console.log('focusNextElement -1', event.keyCode)
                focusNextElement(-1);
    			// Implement your logic to calculate the next index when pressing Up
				break
            case 9: // Tab
                if (event.shiftKey) {
                    console.log('focusNextElement -1', event.keyCode)
                    focusNextElement(-1);
                } else {
                    console.log('focusNextElement 1', event.keyCode)
                    focusNextElement(1);
                }
                break;
			case 39: // Right
			case 40: // Down
                console.log('focusNextElement 1', event.keyCode)
                focusNextElement(1);
				break
			case 13: // Enter
				// Implement your logic for the "Enter" button
                // alert('you selected: ' + focusableArray[currentIndex].innerHTML)
                if (currentItem) {
                    console.log('you selected: ' + currentItem.innerHTML)
                    currentItem.click();
                }
				return // Return to avoid changing focus
		}
	})

    function focusNextElement(direction: number) {
		// Get a NodeList of all focusable elements on the page
        const focusableElements = document.querySelectorAll('.item'); 
        const focusableArray = Array.from(focusableElements); // Convert NodeList to Array

        const curr: any = focusableArray[currentIndex];
        if (curr) {
            curr.style.border = '';
        }

        currentIndex+=direction;
        if (currentIndex > focusableArray.length - 1) {
            currentIndex = 0;
        } else if (currentIndex < 0) {
            currentIndex = focusableArray.length - 1;
        }
        
        (focusableArray[currentIndex] as any).style.border = '2px solid blue';
        (focusableArray[currentIndex] as any).scrollIntoView({behavior: "smooth", block: "center", inline: "nearest"});
        currentItem = focusableArray[currentIndex];

		// get the item currently in focus
		const activeElement: any = document.activeElement;
		if (activeElement) console.log('activeElement', activeElement)
		if (activeElement) activeElement.blur()
		document.body.focus()
    }


</script>
Version: 0.0.2<br/>

<ion-card button={true} class="item i1">
	<ion-card-header>
		<ion-card-title>Card 1</ion-card-title>
		<ion-card-subtitle>Card Subtitle</ion-card-subtitle>
	</ion-card-header>

	<ion-card-content>
		Here's a small text description for the card content. Nothing more, nothing less.
	</ion-card-content>
</ion-card>

<ion-card button={true} class="item i2" on:click={()=>{alert('click')}}>
	<ion-card-header>
		<ion-card-title>Card 2</ion-card-title>
		<ion-card-subtitle>Card Subtitle</ion-card-subtitle>
	</ion-card-header>

	<ion-card-content>
		Here's a small text description for the card content. Nothing more, nothing less.
	</ion-card-content>
</ion-card>

<ion-card button={true} class="item i3">
	<ion-card-header>
		<ion-card-title>Card 3</ion-card-title>
		<ion-card-subtitle>Card Subtitle</ion-card-subtitle>
	</ion-card-header>

	<ion-card-content>
		Here's a small text description for the card content. Nothing more, nothing less.
	</ion-card-content>
</ion-card>

<ion-card button={true} class="item i4">
	<ion-card-header>
		<ion-card-title>Card 4</ion-card-title>
		<ion-card-subtitle>Card Subtitle</ion-card-subtitle>
	</ion-card-header>

	<ion-card-content>
		Here's a small text description for the card content. Nothing more, nothing less.
	</ion-card-content>
</ion-card>

<style>
	ion-card {
		max-width: 400px;
        min-height: 200px;
	}
	ion-card:focus {
        background-color: pink;
	}
</style>
