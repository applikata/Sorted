function sendNumbers() {
    var numbers = document.getElementById('numberInput').value;
    
    // Передача массива чисел на бэкенд
    var url = 'http://localhost:8550/sorted?numbers=' + numbers;
    fetch(url)
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        console.log('Ответ от бэкенда:', data);

        // Пример отображения данных на странице
        var resultElement = document.getElementById('result');
        resultElement.innerHTML = 'Данные отсортированны: ' + JSON.stringify(data);
      })
      .catch(error => {
        console.error('Произошла ошибка:', error);
        // Обработка ошибки
      });
}