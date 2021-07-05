# Observer pattern | Golang

POV: ARMY wants to buy RJ

Using the observer pattern, *armyFirst* and *armySecond* can find out if *"BT21 RJ pillow"* has appeared in stock.

When ARMY want to know if any BT21 toy has appeared in stock, they must use the *subscribe()* method . 
Then when an update happens and the toy the ARMY want to know about is available for sale, ARMY will receive an email with a notification.
If ARMY have already bought a toy (lucky ones) or don't want to know  about the availability of the toy anymore, then they need to use the *unsubscribe()* method.
