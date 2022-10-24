### Select
The `select` statement is the glue that binds channels together; it's how we're able to compose channels together in a program to form larger abstractions.

Just like a `switch` block, a select block encompasses a series of case statements that guard a series of statements; however, that's where similarities end. Unlike switch blocks, case statements in a select block aren't tested sequentially, and execution won't automatically fall through if none of the criteria are met. Instead, all channel reads and writes are considered simultaneously to see if any of them are read. (if none of the channels are ready, the entire select statement blocks.)

#### What happens when multiple channels have something to read?
The Go runtime will perform a pseudorandom uniform selection over the set of case statements. This just means that of your set of case statements, each has an equal chance of being selected as all the others.

The Go runtime cannot know anything about the intent of your select statement. Because of this, the best thing the Go runtime can hope to do is to work well in the average case.

#### What happens if there are never any channels that become read?
If there is nothing useful you can do when all the channels are blocked, but you also can't block forever, you may want to timeout.

#### What if we want to do something but no channels are currently ready?
The select statement also allows for a `default` clause in case you'de like to do something if all the channels you're selecting against are blocking.
Usually you'll see a default clause used in conjuction with a for-select loop. This allows a goroutine to make progress on work while waiting for another goroutine to report a result.

Finally, there is a special case for empty select statements: select statements with no case clauses. `select {}` will simply block forever.