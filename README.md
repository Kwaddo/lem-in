# Lem-in Sorting Algorithm

Made by Ali Khalaf, Salman Nader, Ahmed Aburowais, and Moataz Ibrahim. This complex program is meant to allow the sortings of ants within a certain grid. By utilizing Depth-First Search and many other functions, it is able to give a set order of the paths that ants go through without overlapping over each other.

## How-To-Use

The entire program runs with Go, so with a given IDE you'd want to run the code by typing the following in the terminal:
```
go run main.go [INPUTFILENAME.TXT]
```

There is already a text file called "input.txt" within the folder, but you can use any text file within it. Once you create the text file, the inputs must be according to a strict format. The format is as follows:
```
number_of_ants
the_rooms
the_links

Lx-y
Lx-y Lx-y
Lx-y Lx-y Lx-y
Lx-y Lx-y
Lx-y
```

To be put into practice:
```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```

## Code Explained

The idea of the parse is a checker, to see if there are any conflicting connections, any wrong values, any repetitions, and any formats that completely go against what should follow through.

First thing this is set out to make is an array of paths, so that every validated and confirmed path is thrown into it and used accordingly for the antmove function. The function FindPaths focuses on finding every possible path, and within it validates the paths if there are three or mnore possible paths to begin with since there doesn't need to be any validation if there are only two possible paths.

By using DFS, the given function below, it focuses mainly on returning every possible path with no discrimination. Within it is a function that returns the distance between the two rooms, which is in accordance to pythogoreas's theorem. Once all paths are found, they are to be validated.

The ValidatePaths function is a slightly complex algorithm that only keeps paths that do not overlap each other, keeping each of them unique. "uPaths" is the Paths struct array that has only the ones that are chosen. "roomUsedCount" is a map that is meant to keep the rooms that are used in their specific places, for rememberance that it does not overlap in the future. It checks how many times they are used and keeps a count. A for-loop is made for the paths, to check each path accordingly one-by-one. 

Within the for-loop is the overlap score, which has the total from "roomUsedCount" so that it can be used soon. The "overlapRatio" focuses on dividing the overlap score by the total amount of internal rooms, obviously disregarding the start and end rooms since they do not count. After constant trial and error, we found the best ratio for getting only the unique paths to be smaller than or equal to 0.3. This is because within that given path, there are a small amount of overlaps and by making the number bigger we are only allowing for more overlaps, so 0.3 is meant to be that sweet spot. Within the if statement, it adds the path to "uPaths" and then keeps another for loop to increase the "roomUsedCount". Once all of it is done, it returns the validated paths only. This is not where it ends, as the function returns another called CullPaths that gets rid of any conflicting rooms.

After creating a new Paths array, a for-range is made for the inputted paths and within it is a boolean called "shouldAdd" which should indicate if the chosen path should be added or not. The for-range goes through all paths that are found within. Another for-range is made for the culled paths' variable, which starts empty so it instantly goes to the if statement under it and appends the first path regardless. Now, the for-range is useable, so it first creates a new variable called "minLength" that first is made to be the length of the rooms for the paths. Then, an if statement that checks if the length of the rooms of the chosen culled paths is smaller than the length of the chosen path. When the if statement follows through, minLength is changed to be the length of the chosen culled paths in the current for-range. Another for-loop is made, one to compare the rooms of the culled paths and the normal paths. This is to find any conflicting paths, disregarding the first and last room which is why this loop starts at 1 and ends at minLength minus one. Once one is found within the if statement, it will then change the shouldAdd boolean to false and the function will skip adding the path to the culled paths. After all of that is done, it maintains the paths that have not been culled. 

Within the for-loop is the overlap score, which has the total from "roomUsedCount" so that it can be used soon. The "overlapRatio" focuses on dividing the overlap score by the total amount of internal rooms, obviously disregarding the start and end rooms since they do not count. After constant trial and error, we found the best ratio for getting only the unique paths to be smaller than or equal to 0.3. This is because within that given path, there are a small amount of overlaps and by making the number bigger we are only allowing for more overlaps, so 0.3 is meant to be that sweet spot. Within the if statement, it adds the path to "uPaths" and then keeps another for loop to increase the "roomUsedCount". Once all of it is done, it returns the validated paths only.

After using the DFS algorithm to receive every possible path and using the Validation algorithm to only keep paths that aren't overlapping, we reach this point. The MoveAnts function uses the amount of ants and the given paths as inputs, and then creates four array variables according to the ant amounts which focus on the ant's positions, their paths, and to check if the ant finished or entered the first room.

The first for loop is in accordance to the amount of ants, and the idea of it is that it assigns each ant to a certain path through round robbin style. Our DFS algorithm already sorts the paths from shortest to longest from their lenghts, so ideally we would like the first ant to start from the shortest path, and the next to go to the second-shortest, and so on. There is an extra feature, however, which focuses on assigning the last ant to the shortest path regardless of the round robbin method, only if there are only two paths.

After that, we reach the movement section of the function. A for loop is created that is only done when all ants have their paths set out for them. The movement is just a string array, while a variable called occupied is meant to check if the potential room is occupied by a different ant. After that, a nested loop is created as a normal for-i loop, where "i" here would be the ID of each ant. First thing it checks is if "finishedchecker" is complete, this is because the for loop repeats on every ant but the only way it moves on to the next upcoming one is if "finishedchecker" returns true. After that, it begins taking the ants through their chosen paths.

A new variable named pathID is named such because it is the path in accordance to the current chosen ant, from their ID. Then comes an if statement which checks if the positions, which would be an int and would intentionally increment every single time the given ant follows through a move, for the ID of the ant is smaller than the amount of rooms that the chosen path is. If this returns false, then the ant is considered finished and would move on to the next ant.

Within the previous if statement, a new variable is created named "nextRoom" which is focuses on returning the next room only, since the statements end before the final room. "enteredFirstRoom" is defaulted as false, and the if statement next checks if this current ant has entered the first room. If it did not, the if statement then follows through to another one which checks if the next room is occupied or not. If it isn't occupied, the following happens:
1) The position of the ants increment, assuming it's next trajectory is the next room.
2) A new variable called "currentRoom" is created, and it returns the current ant's room rather than the next one done previously.
3) Since we want to show the ant going through the current movement, it appends the style (L[ID]-[ROOM]) to the "movement" array.
4) Since it is the first room and we want to show that it is occupied, we keep them as true in accordance to the occupied map-boolean and "enteredFirstRoom".

That is all if the first room was entered and if the next room wasn't occupied, if it were it would have skipped it and if the first room wasn't entered then it would assume that it is entering for the first time through the first room, and would do all the steps above disregarding the fourth one. In the end, if there are any inputs in the "movement" array, then it would be printed in according to this until we reach the final row.
