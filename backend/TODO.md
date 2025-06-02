# TODO List for Backend 

## Primitives

Several "primitive" elements are required for the basic logic and operation of the music Player

### Queue

The queue is a list of AudioFiles and handles logic for the following instances:
- add to the queue
- delete from the queue
- enqueue
- dequeue
- selecting the previous/next item
- clear queue
- shuffle queue

The implementation will be an array with circular indexing (for now)

### AudioFile

A struct acting as an abstract datatype for each audiofile. It should contain:
- name
- path
- duration

## Player

The Player is the actual application that manages the Queue and handles the following objects:
- volume
- format
- current song
- song finish
- song start
- song skip
- song control (pause/play)

Player will need methods to:
- set/update the volume
- play/pause a song
- skip a song

For this, the beep library will need to be used at github.com/faiface/beep
