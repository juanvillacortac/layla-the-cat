package com.jvillacorta.layla;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Build;
import android.os.Bundle;
import android.view.View;
import android.view.WindowManager;

import go.Seq;

import com.jvillacorta.layla.layla.Layla;
import com.jvillacorta.layla.layla.EbitenView;

import java.io.File;

public class MainActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        hideSystemUi();
        
        super.onCreate(savedInstanceState);

        Seq.setContext(getApplicationContext());

        File dir = getExternalFilesDir(null);
        Layla.setDataPath(dir.getAbsolutePath());

        setContentView(R.layout.activity_main);
    }

    private void hideSystemUi() {
        View decorView = getWindow().getDecorView();
        decorView.setSystemUiVisibility(
                View.SYSTEM_UI_FLAG_IMMERSIVE_STICKY
                        | View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                        | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION
                        | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                        | View.SYSTEM_UI_FLAG_HIDE_NAVIGATION
                        | View.SYSTEM_UI_FLAG_FULLSCREEN
        );
    }
    private EbitenView getEbitenView() {
        return (EbitenView) this.findViewById(R.id.ebitenview);
    }

    @Override
    protected void onPause() {
        super.onPause();
        this.getEbitenView().suspendGame();
    }

    @Override
    protected void onResume() {
        super.onResume();
        hideSystemUi();
        this.getEbitenView().resumeGame();
    }

    @Override
    public void onBackPressed() {
        Layla.backButton();
    }
}
