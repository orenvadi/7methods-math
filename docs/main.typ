#import "@preview/codelst:1.0.0": sourcecode, sourcefile
#import "template.typ": *

// Take a look at the file `template.typ` in the file panel
// to customize this template and discover how it works.
#show: project.with(
  title: "СРСП №3",
  big_title: "Численные метолы решения нелинейных уравнений",
  authors: (
    "Бактыбеков Нурсултан",
  ),
  // Insert your abstract after the colon, wrapped in brackets.
  // Example: `abstract: [This is my abstract...]`
  date: "2 Октября, 2023",
)



#pagebreak()

#outline(
  indent: 15pt,
  depth: 3
)
#pagebreak()


== Цель работы



Научиться выполнять вычисления с приближенными числами с учетом погрешностей.

== Задание
- Усвоить понятия абсолютной и относительной погрешности и их границ;
- Научиться определять верные, сомнительные, значащие цифры в приближенном числе и определять по ним погрешности приближений;
- Научиться находить погрешности вычислений;
- Вычислить корни уравнения $x^5 - x + 0.2 = 0$

== Выполнение заданий



== График функции уравнения


#figure(
  image("./images/plot.png", width: 80%),
)


#pagebreak()
== Метод деления пополам



=== Решение


#sourcefile(read("../methods/halfDivMethod.go"), file:"../methods/halfDivMethod.go")






#pagebreak()
== Метод Ньютона

=== Решение

#sourcefile(read("../methods/newtonMethod.go"), file:"../methods/newtonMethod.go")







== Метод простой итерации


=== Решение 


#sourcefile(read("../methods/simpleIterationMethod.go"), file:"../methods/simpleIterationMethod.go")






== Метод хорд


=== Решение 


#sourcefile(read("../methods/chordMethod.go"), file:"../methods/chordMethod.go")




== Метод метод секущих


=== Решение


#sourcefile(read("../methods/secantMethod.go"), file:"../methods/secantMethod.go")






#pagebreak()
== Метод Гаусса


=== Решение



#sourcefile(read("../methods/gaussMethod.go"), file:"../methods/gaussMethod.go")




== Метод Крамера


=== Решение


#sourcefile(read("../methods/halfDivMethod.go"), file:"../methods/halfDivMethod.go")







#pagebreak()
== Запуск всех методов

=== Код

#sourcefile(read("../main.go"), file:"../main.go")


=== Результат запуска

#sourcefile(read("../main_output.txt"), file:"../main_output.txt")
