const samples:{ [key:number]:string; } = {
    0: "",
//////// 1: XML
    1: `<?xml version="1.0"?>
<catalog>
   <book id="bk101">
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications
      with XML.</description>
   </book>
   <book id="bk102">
      <author>Ralls, Kim</author>
      <title>Midnight Rain</title>
      <genre>Fantasy</genre>
      <price>5.95</price>
      <publish_date>2000-12-16</publish_date>
      <description>A former architect battles corporate zombies,
      an evil sorceress, and her own childhood to become queen
      of the world.</description>
   </book>
</catalog>
`,
//////// 2: GraphQL
    2: `query {
  getUsers(skip: 0, limit: 5) {
    id
    name
  }
}`,
//////// 3: XML
    3: `/* Hello World program */

main()
{
printf("Hello World");

}
`,
//////// 4: JSON
    4: `{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"GlossSee": "markup",
                    "GlossSeeAlso": ["GML", "XML"]
                }
            }
        }
    }
}
`,
}

export default samples;
