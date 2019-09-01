'use strict';

QUnit.module('Тестируем функцию anagram', function () {
	QUnit.test('Функция работает правильно', function (assert) {
		const input = [
			'кот', 'пила', 'барокко',
			'стоп', 'ток', 'кошка',
			'липа', 'коробка', 'пост'
		];

		const output = [
			[ 'барокко', 'коробка' ],
			[ 'кот', 'ток' ],
			[ 'липа', 'пила' ],
			[ 'пост', 'стоп' ]
		];

		assert.deepEqual(anagram(input), output);
	});

	QUnit.test('Функция работает правильно', function (assert) {
		const input = [
			'ирак', 'кот', 'пила', 'логика',
			'стоп', 'ток', 'кошка',
			'липа', 'коробка', 'пост',
			'иголка', 'раки', 'каир'
		];

		const output = [
			['иголка', 'логика'],
			['ирак', 'каир', 'раки'],
			[ 'кот', 'ток' ],
			[ 'липа', 'пила' ],
			[ 'пост', 'стоп' ]
		];

		assert.deepEqual(anagram(input), output);
	});

	QUnit.test('Функция работает правильно', function (assert) {
		const input = [
			'ирак', 'кот', 'пила'
		];

		const output = [

		];

		assert.deepEqual(anagram(input), output);
	});

	QUnit.test('Функция работает правильно', function (assert) {
		const input = [

		];

		const output = [

		];

		assert.deepEqual(anagram(input), output);
	});

	QUnit.test('Функция работает правильно', function (assert) {
		const input = [
			'ток', 'кот', 'пила', 'отк'
		];

		const output = [

			[ 'кот', 'отк' ,'ток' ]
		];

		assert.deepEqual(anagram(input), output);
	});


});
