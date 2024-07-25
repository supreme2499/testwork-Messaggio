package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

//это обёртка для логов на основе логруса написанная каким-то челом с ютуба,
//а я просто переписал её. её можно использовать в других проектах

// создаём хук он нам нужен для записи логов в консоль и в файл. он принимает в семя враётер - это
// массив врайтов который принимает в себя массив байт и возвращает целое число и логруслевел - это
// массив уровней логирования
type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// метод для хука который просто возвращает уровень логгирования
func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// метой файр, который в себа принимает ентру(метод записи) и может вернуть ошибку
func (hook *writeHook) Fire(entry *logrus.Entry) error {
	//передаём в лаён энтри стринг - запись строки
	line, err := entry.String()
	if err != nil {
		return err
	}
	//перебираем наш массив врайтеров, передав в него строку на которой вызывается лог
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

var e *logrus.Entry

// создаём структуру логгер, которая полностью повтораяет логрус энтри
type Logger struct {
	*logrus.Entry
}

// метод позволяющий нам получить наш логгер в другом файле
func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

// это функция инициализации, она запускается автоматически и её задача это задать фармат логов.
// затем мы создаём папку logs с правами доступа "0644". хатемм мы создаём и открываем файл all.log
// в него мы записываем все логи.
func init() {
	//создаём наш логрус
	l := logrus.New()

	//говорим что мы хотим что бы в логах отображалась номер строки где сработал лог.
	l.SetReportCaller(true)

	//задаём формат для вывода логов и записи в файл
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: true,
		FullTimestamp: true,
	}
	//пытаемся создать папку logs с правами доступа 0644
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}
	//создаём файл all.log в который будет вестись запись всех логов и задаём права доступа
	//0640
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	//отключаем вывод
	l.SetOutput(io.Discard)

	//вызываем хук
	l.AddHook(&writeHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	//задаём уровень логирования
	l.SetLevel(logrus.TraceLevel)

	//инициализируем entry который мы написали ранее
	e = logrus.NewEntry(l)
}
