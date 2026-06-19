[1m<root>[m

[1m[94mmodule[m src/base-elem.js[m

[1m<base-elem>[m

[1mAttributes[m

 [1mName[m           │ [1mSummary[m             
────────────────┼─────────────────────
 [36minherited-attr[m │ Inherited attribute 
[1mSlots[m

 [1mName[m           │ [1mSummary[m        
────────────────┼────────────────
 [36minherited-slot[m │ Inherited slot 
[1mEvents[m

 [1mName[m            │ [1mSummary[m         
─────────────────┼─────────────────
 [36minherited-event[m │ Inherited event 
[1mCSS Properties[m

 [1mName[m              │ [1mSyntax[m  │ [1mSummary[m         
───────────────────┼─────────┼─────────────────
 [36m--inherited-color[m │ [36m<color>[m │ Inherited color 
[1mCSS Parts[m

 [1mName[m           │ [1mSummary[m        
────────────────┼────────────────
 [36minherited-part[m │ Inherited part 


[1m[94mmodule[m src/test-elem.js[m

[1m<test-elem>[m

[1mAttributes[m

 [1mName[m           │ [1mDOM Property[m │ [1mReflects[m │ [1mSummary[m             │ [1mDefault[m │ [1mType[m   │ [1mInherited From[m 
────────────────┼──────────────┼──────────┼─────────────────────┼─────────┼────────┼────────────────
 [36minherited-attr[m │ [36m[m             │          │ Inherited attribute │ [36m[m        │ [36m[m       │ BaseElem       
 [36mfoo[m            │ [36mfoo[m          │ ✅       │ Foo summary         │ [36mbar[m     │ [36mstring[m │                
 [36mbar[m            │ [36mbar[m          │          │                     │ [36m[m        │ [36m[m       │                
[1mSlots[m

 [1mName[m           │ [1mSummary[m           │ [1mInherited From[m 
────────────────┼───────────────────┼────────────────
 [36minherited-slot[m │ Inherited slot    │ BaseElem       
 [36m<default>[m      │ Default slot      │                
 [36micon[m           │ Slot for the icon │                
[1mEvents[m

 [1mName[m            │ [1mType[m                     │ [1mSummary[m         │ [1mInherited From[m 
─────────────────┼──────────────────────────┼─────────────────┼────────────────
 [36minherited-event[m │ [36m[m                         │ Inherited event │ BaseElem       
 [36msubmitted[m       │ [36mCustomEvent<'a'|'b'|'c'>[m │ Fires on submit │                
 [36mreset[m           │ [36m[m                         │                 │                
[1mFields[m

 [1mName[m │ [1mSummary[m           
──────┼───────────────────
 [36mfoo[m  │ Foo field summary 
 [36mbar[m  │                   
[1mMethods[m

 [1mName[m         │ [1mReturn Type[m          │ [1mPrivacy[m │ [1mStatic[m │ [1mSummary[m      
──────────────┼──────────────────────┼─────────┼────────┼──────────────
 [36mdoThing[m      │ [36mvoid | Promise<void>[m │ public  │ [36mfalse[m  │ Does a thing 
 [36mprivateStuff[m │ [36mstring[m               │ private │ [36mtrue[m   │              
[1mCSS Properties[m

 [1mName[m              │ [1mSyntax[m  │ [1mDefault[m │ [1mSummary[m         │ [1mInherited From[m 
───────────────────┼─────────┼─────────┼─────────────────┼────────────────
 [36m--inherited-color[m │ [36m<color>[m │ [36m[m        │ Inherited color │ BaseElem       
 [36m--main-color[m      │ [36m<color>[m │ [36m#fff[m    │ Main color      │                
 [36m--unused[m          │ [36m[m        │ [36m[m        │                 │                
[1mCSS Parts[m

 [1mName[m           │ [1mSummary[m        │ [1mInherited From[m 
────────────────┼────────────────┼────────────────
 [36minherited-part[m │ Inherited part │ BaseElem       
 [36mlabel[m          │ The label part │                
[1mCSS States[m

 [1mName[m       │ [1mSummary[m      
────────────┼──────────────
 [36m--active[m   │ Active state 
 [36m--inactive[m │              

