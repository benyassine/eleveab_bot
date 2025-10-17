# أرشيف الدراسة — موقع Hugo جاهز للنشر على Netlify

موقع تعليمي عربي باتجاه RTL وخط Cairo، مع قالب بسيط وحديث، جاهز للنشر على GitHub وNetlify.

## المتطلبات
- تثبيت Hugo (Extended)
- Git

## التشغيل محليًا
```bash
hugo server
```
ثم افتح المتصفح على: http://localhost:1313

## البنية
- `themes/arshif-aldirasa/` القالب المخصص
- `content/` المحتوى (القوانين، الدروس، الأخبار، حول)
- `assets/` ملفات CSS
- `static/` ملفات ثابتة (PDF، أيقونات)

## الإعداد
- تعديل `config.toml` لتغيير `title` و`description`
- لإضافة Google Analytics ضع معرف GA4 في:
  ```toml
  [services.googleAnalytics]
  id = "G-XXXXXXXXXX"
  ```

## النشر على GitHub
1. أنشئ مستودعًا جديدًا على GitHub.
2. ادفع الشفرة:
   ```bash
   git init
   git add .
   git commit -m "إطلاق: أرشيف الدراسة"
   git branch -M main
   git remote add origin git@github.com:USER/REPO.git
   git push -u origin main
   ```

## الربط مع Netlify للنشر التلقائي
1. ادخل إلى Netlify واربط المستودع.
2. في الإعدادات سيكتشف Netlify الملف `netlify.toml` تلقائيًا:
   - Build command: `hugo --gc --minify`
   - Publish directory: `public`
3. اضغط Deploy.

## محتوى تجريبي
- درس: "الدرس الأول في الفيزياء: الطاقة" مع ملف PDF
- قانون: "القانون التوجيهي للتعليم لسنة 2025"
- خبر: "تحديث رزنامة الامتحانات الوطنية"
- صفحة "حول الموقع"

## ملاحظات
- اتجاه الموقع RTL بشكل كامل.
- مشاركات اجتماعية (فيسبوك، تويتر، واتساب) متاحة في الصفحات الفردية.
