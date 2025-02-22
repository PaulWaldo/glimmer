# Architecting the Users Groups Photos page

## Goal

The goal is to obtain and display photos from all of the Flickr groups to which the user belongs.
Since the user may belong to a number of groups, the main organization will be these groups.  Each group is to be displayed as a series of cards in a scrollable pane.  The title of this card will be the group name.
Within each group card, X number of photos will be displayed in a grid, where X is a configurable parameter.
At the bottom of each group card will be a tappable object titled "More...".  When this object is tapped, additional photos from the group will be displayed.  There will also be a UI element that will allow the group to be collapsed.  There will still be a card with the group name, but the photos display will be collapsed.

## Requirements

* All operations need to run in the background, to not block smooth operation of the UI
* TDD must be followed
* If anything is unclear, you must ask before continuing

## Implementation Considerations

* Group photo information can be obtained with the `api.GetUsersGroupPhotos` function
* The UI must be kept responsive
* Encapsulation and extensibility are key
* The number of photos can be quite large, so downloading them should be done in batches.  The batch size should be configurable.

## General Process Flow

1. The UI is constructed with an AppTab; one tab for the contact photos and the other for group photos.  The goup photos pane is an empty grid view to start.
2. The list of group photos is obtained.  This list contains URLs to the online images.
3. Wait for the main run loop to be fully started so the UI is available.
4. For all the groups listed in the photo list, create and display a card for that group
5. For each group, loop through the listed photos and create a photo card.  The card will be responsible for taking the information about the photo and displaying it properly, including
    * Title
    * Author
    * Image as provided by the URL.  The card is responsible for downloading this image content
    The number of simultaneous image download should be configurable
6. Place all the photo cards for a particular group inside the group card
7. If the "More..." button is tapped, create another batch of photocards for that group and add them to the group card
8. If the "Collapse" element is tapped, accordian the group card so no photos are displayed.  An "expand" element should be visible to later re-show the photo cards.

## Our process

1. You will first digest this document
2. If there is anything that is unclear, you must ask me and we will discuss it
3. Once we both agree that we fully understand the plan, you will create a series of user stories.  These stories will break down the big problem into smaller step-by-step tasks.
4. You will modify this document in the `User Stories` section to provide
    * a summary of the overall plan
    * a list of detailed user stories that decompose the entire problem into managebale chunks
    * a checkbox for each story that we will use as we work to indicate what we have completed and what work is left to do.  This will enable us to take a break and come back later without forgetting what we have done.

## User Stories
