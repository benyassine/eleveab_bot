# أرشف الدارسة - موقع تعليمي عربي

موقع تعليمي عربي شامل مبني باستخدام Hugo، مصمم خصيصاً للمحتوى التعليمي باللغة العربية مع دعم كامل للاتجاه من اليمين إلى اليسار (RTL).

## 🌟 المميزات

- **دعم RTL كامل**: مصمم خصيصاً للغة العربية
- **تصميم متجاوب**: يعمل على جميع الأجهزة والشاشات
- **خط Cairo**: خط عربي جميل وواضح
- **تحسين SEO**: محسن لمحركات البحث
- **مشاركة اجتماعية**: أزرار مشاركة لفيسبوك وتويتر وواتساب
- **عرض PDF**: إمكانية عرض ملفات PDF داخل الصفحات
- **جاهز للنشر**: مكون للعمل مع Netlify وGitHub

## 📁 هيكل المشروع

```
arshif-aldirasa-site/
├── archetypes/          # قوالب المحتوى
├── content/             # المحتوى
│   ├── laws/           # القوانين التعليمية
│   ├── lessons/        # الدروس والمذكرات
│   ├── news/           # الأخبار التربوية
│   └── about/          # صفحة حول الموقع
├── data/               # ملفات البيانات
├── layouts/            # قوالب HTML
├── static/             # الملفات الثابتة
│   ├── css/           # ملفات CSS
│   ├── js/            # ملفات JavaScript
│   ├── images/        # الصور
│   └── pdfs/          # ملفات PDF
├── themes/             # القالب
│   └── arshif-aldirasa/
├── config.toml         # إعدادات الموقع
├── netlify.toml        # إعدادات Netlify
└── README.md           # هذا الملف
```

## 🚀 التشغيل المحلي

### المتطلبات

- [Hugo](https://gohugo.io/) (الإصدار 0.80.0 أو أحدث)
- متصفح ويب حديث

### خطوات التشغيل

1. **تحميل Hugo**:
   ```bash
   # على Ubuntu/Debian
   sudo apt-get install hugo
   
   # على macOS
   brew install hugo
   
   # على Windows
   choco install hugo
   ```

2. **استنساخ المشروع**:
   ```bash
   git clone https://github.com/your-username/arshif-aldirasa-site.git
   cd arshif-aldirasa-site
   ```

3. **تشغيل الخادم المحلي**:
   ```bash
   hugo server
   ```

4. **فتح الموقع**:
   افتح المتصفح وانتقل إلى `http://localhost:1313`

## 📝 إضافة محتوى جديد

### إضافة درس جديد

```bash
hugo new lessons/اسم-الدرس.md
```

### إضافة قانون جديد

```bash
hugo new laws/اسم-القانون.md
```

### إضافة خبر جديد

```bash
hugo new news/اسم-الخبر.md
```

## 🎨 تخصيص التصميم

### تغيير الألوان

قم بتعديل المتغيرات في `themes/arshif-aldirasa/static/css/main.css`:

```css
:root {
  --primary-color: #0077b6;    /* اللون الأساسي */
  --background-color: #ffffff; /* لون الخلفية */
  --text-color: #333;          /* لون النص */
}
```

### تغيير الخط

قم بتعديل `fontFamily` في `config.toml`:

```toml
[params]
  fontFamily = "Cairo, Arial, sans-serif"
```

## 🔧 الإعدادات

### إعداد Google Analytics

قم بتعديل `config.toml`:

```toml
[params]
  googleAnalytics = "GA_TRACKING_ID"
```

### إعداد وسائل التواصل الاجتماعي

```toml
[params]
  twitter = "@your_twitter"
  facebook = "your.facebook.page"
```

## 📤 النشر على Netlify

### الطريقة الأولى: رفع مباشر

1. **ضغط المشروع**:
   ```bash
   zip -r arshif-aldirasa-site.zip .
   ```

2. **رفع إلى Netlify**:
   - اذهب إلى [netlify.com](https://netlify.com)
   - سجل دخول أو أنشئ حساب
   - اسحب وأفلت ملف ZIP في منطقة "Deploy manually"

### الطريقة الثانية: ربط GitHub

1. **رفع المشروع إلى GitHub**:
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/your-username/arshif-aldirasa-site.git
   git push -u origin main
   ```

2. **ربط Netlify بـ GitHub**:
   - اذهب إلى Netlify
   - اختر "New site from Git"
   - اختر GitHub وحدد المستودع
   - اضبط إعدادات البناء:
     - Build command: `hugo --minify`
     - Publish directory: `public`

## 📱 الميزات المتقدمة

### عرض ملفات PDF

```markdown
---
title: "عنوان الدرس"
pdf: "/pdfs/lesson.pdf"
---
```

### مشاركة اجتماعية

```markdown
---
title: "عنوان المقال"
share: true
---
```

### إضافة صور

```markdown
![وصف الصورة](/images/image.jpg)
```

## 🛠️ استكشاف الأخطاء

### مشاكل شائعة

1. **الموقع لا يظهر**:
   - تأكد من تشغيل `hugo server`
   - تحقق من أن المنفذ 1313 متاح

2. **التصميم لا يظهر**:
   - تأكد من وجود ملف CSS في `static/css/`
   - تحقق من المسار في `head.html`

3. **مشاكل RTL**:
   - تأكد من `dir="rtl"` في HTML
   - تحقق من إعدادات CSS للاتجاه

## 📞 الدعم

للحصول على المساعدة:

- **البريد الإلكتروني**: support@arshif-aldirasa.com
- **GitHub Issues**: [إنشاء مشكلة جديدة](https://github.com/your-username/arshif-aldirasa-site/issues)
- **التوثيق**: [Hugo Documentation](https://gohugo.io/documentation/)

## 📄 الترخيص

هذا المشروع مرخص تحت رخصة MIT. راجع ملف [LICENSE](LICENSE) للتفاصيل.

## 🤝 المساهمة

نرحب بمساهماتكم! يرجى:

1. عمل Fork للمشروع
2. إنشاء فرع للميزة الجديدة (`git checkout -b feature/amazing-feature`)
3. عمل Commit للتغييرات (`git commit -m 'Add amazing feature'`)
4. دفع الفرع (`git push origin feature/amazing-feature`)
5. فتح Pull Request

## 📈 التطوير المستقبلي

- [ ] إضافة نظام تعليقات
- [ ] دعم الفيديو والصوت
- [ ] تطبيق محمول
- [ ] نظام تسجيل الدخول
- [ ] لوحة تحكم إدارية

---

**أرشف الدارسة** - نافذتك على عالم التعليم العربي

*تم إنشاؤه بـ ❤️ للمجتمع التعليمي العربي*