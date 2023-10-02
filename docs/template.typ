// The project function defines how your document looks.
// It takes your content and some metadata and formats it.
// Go ahead and customize it to your liking!

#let project(
  title: "",
  big_title: "",
  authors: (),
  date: none,
  body,
) = {
  // Set the document's basic properties.
  set document(author: authors, title: title)
  set page(
    margin: (left: 20mm, right: 20mm, top: 25mm, bottom: 25mm),
    numbering: "1",
    number-align: center,
  )

  // Save heading and body font families in variables.
  let body-font = "Minion Pro"
  // let sans-font = "Petersburg"
  // let sans-font = "IBM Plex Sans"
  let sans-font = "Cantarell"

  // Set body font family.
  set text(font: body-font, lang: "ru", size: 12pt)
  show heading: set text(font: sans-font)


  show outline.entry.where(
  level: 2
): it => {
  v(12pt, weak: true)
  strong(it)
}

  show link: underline
  

  align(center)[
    #v(15em)
  ]
  // Title row.
  align(center)[
    #block(text(font: sans-font, weight: 700, 1.75em, title))
    #v(2em, weak: true)
  ]

  align(center)[
    #block(text(font: sans-font, weight: 700, 1.55em, big_title))
    #v(1em, weak: true)
  ]

  // Author information.
  pad(
    top: 0.5em,
    bottom: 0.5em,
    x: 2em,
    grid(
      columns: (1fr,) * calc.min(3, authors.len()),
      gutter: 1em,
      ..authors.map(author => align(center, strong(author))),
    ),
  )
  align(center)[
    #date
  ]
  

  show raw.where(block: true): txt => pad(
      left: 0.5em,
      block(
        width: 100%,
        radius: 0.5em,
        //stroke: luma(230),
        fill: luma(240),
        pad(
          left: 1em,
          right: 1em,
          top: 1em,
          bottom: 1em,
          txt
        )
      )
    )

//   raw((
//         width: 100%,
//         radius: 0.5em,
//         //stroke: luma(230),
//         fill: luma(240),
//         pad(
//           left: 1em,
//           right: 1em,
//           top: 1em,
//           bottom: 1em,
//           txt
//         ),
// )
  // Main body.
  set par(justify: true)


  body
}
