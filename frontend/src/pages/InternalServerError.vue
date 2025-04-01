<template>
  <div class="errorContainer">
    <div class="errorContainer__inner error">
      <h2 class="error__title">500</h2>
      <p class="error__name">Internal Server Error</p>
      <p class="error__detail">
        アクセスしようとしたページは<br />
        表示できませんでした。<br />
        お手数ですが、しばらく時間を置いてから再度お試しください。
      </p>
      <div class="pagetopButtonWrap">
        <RouterLink to="/" class="pagetopButtonWrap__button">TOPページに戻る</RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.errorContainer {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f8f8f8;
}

.errorContainer__inner {
  text-align: center;
  background-color: #fff;
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.error__title {
  font-size: 96px;
  font-weight: bold;
  color: #e74c3c;
  margin: 0;
}

.error__name {
  font-size: 24px;
  color: #333;
  margin-top: 10px;
}

.error__detail {
  font-size: 16px;
  color: #666;
  margin-top: 20px;
  line-height: 1.6;
}

.pagetopButtonWrap {
  margin-top: 30px;
}

.pagetopButtonWrap__button {
  display: inline-block;
  background-color: #3498db;
  color: white;
  text-decoration: none;
  padding: 12px 24px;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.pagetopButtonWrap__button:hover {
  background-color: #2980b9;
}
</style>
